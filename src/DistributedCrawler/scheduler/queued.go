package scheduler

import (
	"DistributedCrawler/engine"
)

//队列调度器
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

//提交任务
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

//func (s *QueuedScheduler)ConfigureMasterWorkerChan(c chan engine.Request)  {
//}
func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		//建立Request队列和worker队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			//查看request和worker是否同时存在，若有则取出
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			//调度器中有请求时，将请求加入到请求队列
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			//调度器中有可以接收任务的worker时，将请求加入到worker中
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			//当同时有请求又有worker时，将请求分配给worker执行，并从队列中移除
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
