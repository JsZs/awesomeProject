package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
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
	Submit(Request)
	//ConfigureMasterWorkerChan(chan Request)
	//WorkerReady(chan Request)
	WorkChan() chan Request
	Run()
	ReadyNotifier
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r) //scheduler把请求提交到request chan中
	}
	//in:=make(chan Request)
	out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	//分发任务
	e.Scheduler.Run()
	//for i:=0;i<e.WorkerCount;i++{
	//	createWorker(in,out)
	//}
	//创建多个work
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("Got item: %v",item)
			go func() { e.ItemChan <- item }()
		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
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
	go func() {
		for {

			ready.WorkerReady(in) //告诉scheduler,worker准备好了
			request := <-in
			// 换成call rpc, 这里是分布式的关键
			result, e := e.RequestProcessor(request)
			//result, e := Worker(request)
			if e != nil {
				continue
			}
			out <- result
		}
	}()

}
