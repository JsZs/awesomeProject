package parser

import (
	"awesomeProject/src/ConcurrentCrawler/config"
	"awesomeProject/src/DistributedCrawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"
[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		//result.Items=append(result.Items,string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				//ParserFunc:ParseCity,
				Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
			})
	}
	return result
}
