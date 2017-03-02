package main

import (
	"fmt"
	"github.com/cplusgo/go-payment/payment"
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
	fmt.Println(payment.MapToXMLString(datas))
	fmt.Println(payment.MakeSign(datas))
}