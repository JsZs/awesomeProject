package main

import (
	"ConsoleCrawler/engine"
	"ConsoleCrawler/zhenai/parser"
)


func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}
