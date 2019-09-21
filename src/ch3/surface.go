package main

import(
  "fmt"
  "math"
  "strconv"
)

const (
  width, height = 600, 320              // canvas size in pixels
  cells         = 100                   // number of grid cells
  xyrange       = 30.0                  // axis ranges (-xyrange..+xyrange)
  xyscale       = width / 2 / xyrange   // pixels per x or y unit
  zscale        = height * 0.4          // pixels per z unit
  angle         = math.Pi / 6           // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)  // sin(30°), cos(30°)

func main() {
  fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
      "style='stroke: grey; fill: white; stroke-width: 0.7' "+
      "width='%d' height='%d'>", width, height)
  var color string = ""
  var min float64 = -0.22
  for i := 0; i < cells; i++ {
    for j := 0; j < cells; j++ {
      ax, ay := corner(i+1, j)
      bx, by := corner(i, j)
      cx, cy := corner(i, j+1)
      dx, dy := corner(i+1, j+1)
      z := f(xyrange * (float64(i)/cells - 0.5), xyrange * (float64(j)/cells - 0.5))
      if math.IsNaN(z){
        color = getColor(255)
      }else{
        color = getColor(int((z-min)/1.2*255))
      }
      if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy){
        continue
      }
      fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' "+"style='stroke: "+"#"+color+"; fill: white; stroke-width: 0.7' "+"/>\n",
          ax, ay, bx, by, cx, cy, dx, dy)
    }
  }
  fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
  // Find point (x,y) at corner of cell (i,j).
  x := xyrange * (float64(i)/cells - 0.5)
  y := xyrange * (float64(j)/cells - 0.5)

  // Compute surface height z.
  z := f(x, y)

  // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
  sx := width/2 + (x-y)*cos30*xyscale
  sy := height/2 + (x+y)*sin30*xyscale - z*zscale
  return sx, sy
}

func f(x, y float64) float64 {
  r := math.Hypot(x, y) // distance from (0,0)
  return math.Sin(r) / r
}

func getColor(rgb int) string {
  red, blue := 0, 0
  if rgb > 255{
    red, blue = 255, 0
  }else{
    red, blue = rgb, 255-rgb
  }
  r, b := RToX(red), RToX(blue)
  return r+"00"+b
}

func RToX(num int) string {
  var mode []string = []string{"A", "B", "C", "D", "E", "F"};
  var result string
  res := num % 16
  num = num / 16

  if num >= 10{
    result += mode[num-10]
  }else{
    result += strconv.Itoa(num)
  }

  if res >= 10{
    result += mode[res-10]
  }else{
    result += strconv.Itoa(res)
  }


  return result
}
