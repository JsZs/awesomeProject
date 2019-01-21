package parser

import (
	"awesomeProject/src/ConcurrentCrawler/config"
	"awesomeProject/src/DistributedCrawler/engine"
	"regexp"
	//"runtime/pprof"
)

//获取信息格式
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

var (
	profileRe = regexp.MustCompile(cityRe)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

//contents为城市页面地址，从每个城市第一页中筛选信息
func ParseCity(contents []byte, _ string) engine.ParseResult {
	//re:=regexp.MustCompile(cityRe)
	//搜索全部与格式相同的信息，返回二维byte数组
	matches := profileRe.FindAllSubmatch(contents, -1)
	//创建结构体进行存放
	result := engine.ParseResult{}
	//把第一张页面中的用户名及地址取出
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
		//往Requests里续request
		result.Requests = append(result.Requests, engine.Request{

			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return result
}
