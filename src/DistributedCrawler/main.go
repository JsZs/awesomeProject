package DistributedCrawler

import (
	"awesomeProject/src/DistributedCrawler/engine"
	"awesomeProject/src/DistributedCrawler/saver"
	"awesomeProject/src/DistributedCrawler/scheduler"
	"awesomeProject/src/DistributedCrawler/zhenai/parser"
)

func main() {
	itemChan, err := saver.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	//e:=engine.ConcurrentEngine{
	//	Scheduler:&scheduler.SimpleScheduler{},
	//	WorkerCount:10,
	//}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		//ParserFunc:parser.ParseCityList,
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
