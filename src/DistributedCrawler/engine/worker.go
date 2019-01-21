package engine

import (
	"awesomeProject/src/DistributedCrawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	//第一次打印为main函数传入地址，然后每次打印是从r.ParserFunc函数中提取出的城市地址
	log.Printf("Fetching %s", r.Url)
	//将不同URL传输进去，返回不同的页面源代码
	body, err := fetcher.Fetch(r.Url)
	//判断URL是否正确 如果不正确 跳过此次循环
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}
	//return r.ParserFunc(body,nil
	return r.Parser.Parse(body, r.Url), nil
}
