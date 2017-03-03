package main

import (
	"github.com/cplusgo/go-payment/helper"
	"github.com/cplusgo/go-payment/payment"
)

func main() {
	appId := "appId"
	appKey := "appKey"
	mchId := "mchId"
	wxSigned := payment.NewWxPaymentSigned(
		appId,
		appKey,
		mchId,
		helper.MD5("Hello"),
		"Body",
		"Hello",
		100,
		"http://payment.example.com/notify",
		"98",
		"111.121.30.166",
		"APP",
		"ziyun")
	wxSigned.Unifiedorder()
}
