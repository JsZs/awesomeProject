package engine

//集合成员的作用 使Run为具体的那个对象 因为sumple中也定义Run函数 主要作用是区分
type ConcurrentEngine struct {
	//将接口作为子集使用
	Scheduler Scheduler
	//控制goroutine数量
	WorkerCount      int
	ItemChan         chan Item //ItemSaver
	RequestProcessor Processor //RPC会调用
}

type Processor func(request Request) (ParseResult, error)

var visitedUrl = make(map[string]bool) //map判断url是否重复

func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	//ConfigureMasterWorkerChan(chan Request)
	//WorkerReady(chan Request)
	WorkChan() chan Request
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r) //scheduler把请求提交到request chan中
	}
	//创建接收者
	//in:=make(chan Request)
	//创建发送者
	out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	//分发任务
	e.Scheduler.Run()
	//for i:=0;i<e.WorkerCount;i++{
	//	createWorker(in,out)
	//}
	//创建多个work（goroutine）
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	for {
		//chan传输的值
		result := <-out
		//获取result的值
		for _, item := range result.Items {
			//log.Printf("Got item: %v",item)
			go func() { e.ItemChan <- item }()
		}
		//提取地址与正则表达式函数
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			//将新的地址传入函数
			e.Scheduler.Submit(request)
		}
	}
}

/**
每个work开启一个goroutine
*/
//func createWorker(out chan ParseResult,s Scheduler)  {
func (e *ConcurrentEngine) createWorker(
	in chan Request,
	out chan ParseResult, ready ReadyNotifier) {
	//in := make(chan Request)
	//直接goroutine 无限接收数据并传送数据 外围for循环控制几个这样的goroutine同时并发
	go func() {
		for {
			ready.WorkerReady(in) //告诉scheduler,worker准备好了
			//将in的值赋值给request没问题 问题是 同时并发并且无限接收数据的in是从哪里来的
			//注意了 关键的是in有一个指针指向in 这个指针也在goroutine无限循环接收
			request := <-in
			//取得ParseResult返回值
			result, e := e.RequestProcessor(request)
			//result, e := Worker(request)
			//如果报错 跳过循环
			if e != nil {
				continue
			}
			//将返回值放入out chan类型，可以放入无限的数据
			out <- result
		}
	}()

}
