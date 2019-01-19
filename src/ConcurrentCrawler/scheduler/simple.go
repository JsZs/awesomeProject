package scheduler

import (
	"ConcurrentCrawler/engine"
)

type SimpleScheduler struct {
	worker chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()

}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
