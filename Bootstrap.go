package main

import (
	"fmt"
	"github.com/cplusgo/go-payment/helper"
)

type Bootstrap struct {
	
}

func (this *Bootstrap) Start()  {
	
}

func main()  {
	datas := make(map[string]string)
	datas["name"] = "hello"
	datas["age"] = "90"
	datas["huge"] = "huge"
	datas["msg"] = "你的天"
	fmt.Println(helper.MapToXMLString(datas))
	fmt.Println(helper.MakeSign(datas))

	fmt.Println(helper.TimeMd5())
}