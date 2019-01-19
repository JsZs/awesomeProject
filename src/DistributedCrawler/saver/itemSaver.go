package saver

import (
	"DistributedCrawler/engine"
	"context"
	"errors"
	"log"
)

func ItemSaver(index string)(chan engine.Item,err error){
	client,e=elastic.NewClient(elastic.SetSniff(false))

	if e!=nil{
		return nil,e

	}
	out:=make(chan engine.Item)

	go func() {
		itemCount:=0
		for{
			item :<- out
			log.Printf("ItemSaver: got item #%d: %v",itemCount,item)
			itemCount++
			err :=Save(client,index,item)
			if err!=nil{
				log.Printf("Item saver:error saving item %v: %v",item,err)
			}
		}
	}()
	return out,nil
}

func Save(client *elastic.Client,index string,item engine.Item) error{
	if item.Type== ""{

		return errors.New("Must supply type")
	}
	indexService :=client.Index().Index(index).Type(item.Type).BodyJson(item)

	if item.Id !=""{
		indexService.Id(item.Id)
	}

	__ ,err :=indexService.Do(context.Background())

	if err!=nil{
		return err
	}
	return nil
}