package scheduler

import (
	"awesomeProject/src/DistributedCrawler/engine"
)

//采用的队列方法进行chan控制
type QueuedScheduler struct {
	engine.Scheduler
	//接收数据，每一次接收后会马上被存入数组中
	requestChan chan engine.Request
	//接收chan类型进行 chan之间的chan传送 ，简单来说就是用一个chan来传一个chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

//向接收数据的requestChan中存放不同的request
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

//让定义的chanchan类型指向外面in的地址,只要chanchan数据变化,外面in同时变化
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

//func (s *QueuedScheduler)ConfigureMasterWorkerChan(c chan engine.Request)  {
//}
func (s *QueuedScheduler) Run() {
	//创建chanchan类型,只有创建后才能接收chan类型
	s.workerChan = make(chan chan engine.Request)
	//创建chan类型,用于接收数据
	s.requestChan = make(chan engine.Request)

	go func() {
		//建立Request队列和worker队列,先进先出,存取所有的需求
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			//同时满足两个需求,有数据并且有人要数据,则取数据
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			//进行chan判断
			select {
			//调度器中有请求时，将请求加入到请求队列
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			//调度器中有可以接收任务的worker时，将请求加入到worker中
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			//有人要数据，并且我这里有数据，则传送出去
			case activeWorker <- activeRequest:
				//传送成功，需求-1
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
