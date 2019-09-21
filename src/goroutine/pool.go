package main

import(
  "fmt"
  "time"
)
//-------------------有关Task任务角色的功能---------------------------
//定义一个任务类型Task
type Task struct {
  f func () error //一个Task里面应该有一个具体的业务，业务名称就叫f
}

func NewTask(arg_f func() error) *Task {
  t := Task{
    f :arg_f,
  }

  return &t
}

//Task也需要一个执行业务的方法
func (t *Task) Execute(){
  t.f()
}

//-----------------------------有关协程池Pool角色的功能-----------------------------

//定义一个Pool协程池的类型

type Pool struct{
  //对外的Task入口 EntryChannel
  EntryChannel chan *Task

  //内部的Task队列 JobsChannel
  JobsChannel chan *Task

  //协程池中最大的worker的数量
  worker_num int
}

//创建Pool的函数

func NewPool(cap int) *Pool{
  //创建一个Pool
  p := Pool{
    EntryChannel : make(chan *Task),
    JobsChannel : make(chan *Task),
    worker_num : cap,
  }
  //返回这个Pool
  return &p
}

//协程池创建一个Worker，并且让Worker去工作
func (p *Pool) worker(worker_ID int) {
  //worker的具体工作

  //1、永久的从JobsChannel中去任务
  for task := range p.JobsChannel{
    //task就是当前worker从JobsChannel中拿到的任务
    //2、一旦渠道任务，执行任务
    task.Execute()
    fmt.Println("worker_ID:", worker_ID, "执行完了一个任务")
  }
}

//让协程池真正的工作
func (p *Pool) run(){
  //1、根据worker_num来创建worker去工作
  for i:=0; i<p.worker_num; i++{
    go p.worker(i)
  }
  //2、不断从EntryChannel中取Task，将取到的Task发送给JobsChannel
  for task := range p.EntryChannel{
    //一旦有Task读到，把任务交给JobsChannel
    p.JobsChannel <- task
  }
}

//------------------------------主函数，测试协程池的工作---------------------------
func main(){
  //1、创建一些任务
  t := NewTask(func () error {
    //当前任务的业务，打印当前的系统时间
    fmt.Println(time.Now())
    return nil
  })

  //2、创建一个Pool协程池，最大worker数量为4
  p := NewPool(4)

  //3、将这些任务交个协程池Pool
  go func(){
    for i:=1; i <= 100; i++{
      //不断的向p中写入任务t
      p.EntryChannel <- t
    }
  }()

  //4、启动Pool
  p.run()
}
