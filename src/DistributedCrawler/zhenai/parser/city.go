package parser

import (
	"ConcurrentCrawler/config"
	"DistributedCrawler/engine"
	"regexp"
	//"runtime/pprof"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

var (
	profileRe = regexp.MustCompile(cityRe)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	//re:=regexp.MustCompile(cityRe)
	//返回二维byte数组
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		//result.Items=append(result.Items,"User "+name)
		result.Requests = append(
			result.Requests, engine.Request{
				Url: url,
				//ParserFunc: func(c []byte) engine.ParseResult {
				//return ParseProfile(c,name)
				//},
				Parser: NewProfileParser(name),
			})
	}
	//return  result
	matches = cityUrlRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		//往Requests里续equest
		result.Requests = append(result.Requests, engine.Request{

			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return result
}
