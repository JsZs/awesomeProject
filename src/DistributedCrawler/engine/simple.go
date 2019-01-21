package engine

import (
	"log"
)

type SimpleEngine struct {
}

//引擎控制整个程序的流程
func (e SimpleEngine) Run(seeds ...Request) {
	//接收main函数传过来的值
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	//用传过来的值进行解析及提取
	for len(requests) > 0 {
		//获取第一个值
		r := requests[0]
		//进行切片,把已经提取的内容筛选出去
		requests = requests[1:]
		//接收worker函数的值
		parseResult, err := Worker(r)

		if err != nil {
			continue
		}
		//requests被填满,requests又得到新的URL和运算函数,被抓取信息只要足够就可以一直运行下去
		requests = append(requests, parseResult.Requests...)
		//打印所有在PrintCityList函数返回的Item值,Item值是任何类型可以使城市名也可以是用户信息
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
