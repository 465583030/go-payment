package payment

import (
	"strconv"
	"fmt"
)

// 参考官方文档：
// 统一下单接口:https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1
// 调起支付:https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12&index=2

const wxUnifiedorderURL string = "https://api.mch.weixin.qq.com/pay/unifiedorder"

type WxPaymentSigned struct {
	appId      string
	appKey     string
	mchId      string
	nonceStr   string
	body       string
	desc       string
	fee        int
	notifyUrl  string
	outTradeNo string
	createIp   string
	tradeType  string
	attach     string
}

func NewWxPaymentSigned(appId string, appKey string, mchId string, nonceStr string, body string,
	desc string, fee int, notifyUrl string, outTradeNo string, createIp string,
	tradeType string, attach string) *WxPaymentSigned {
	payment := &WxPaymentSigned{
		appId:      appId,
		appKey:     appKey,
		mchId:      mchId,
		nonceStr:   nonceStr,
		body:       body,
		desc:       desc,
		fee:        fee,
		notifyUrl:  notifyUrl,
		outTradeNo: outTradeNo,
		createIp:   createIp,
		tradeType:  tradeType,
		attach:     attach,
	}
	return payment
}

func (this *WxPaymentSigned) unifiedorder() {
	presignData := make(map[string]string)
	presignData["appid"] = this.appId
	presignData["mch_id"] = this.mchId
	presignData["nonce_str"] = this.nonceStr
	presignData["body"] = this.body
	presignData["out_trade_no"] = this.outTradeNo
	presignData["total_fee"] = strconv.Itoa(this.fee)
	presignData["spbill_create_ip"] = this.createIp
	presignData["notify_url"] = this.notifyUrl
	presignData["trade_type"] = this.tradeType
	presignData["attach"] = this.attach

	for k, v := range presignData {
		fmt.Println(fmt.Sprintf("key: %s; value:%s\n", k, v))
	}
}

func mapToXMLString(data map[string]string) string {
	xmlString := "<xml>"
	for key, value := range data {
		xmlString += "<" + key + "><![CDATA[" + value + "]]></" + key + ">"
	}
	xmlString += "</xml>"
	return xmlString
}

func makeSign(params map[string]string) {
	keys := []string{}
	for k, _ := range params {
		keys = append(keys, k)
	}
}

func (this *WxPaymentSigned) Signed() {
	presignData := make(map[string]string)
	presignData["appid"] = this.appId
	presignData["partnerid"] = this.mchId
	presignData["prepayid"] = this.nonceStr
	presignData["package"] = this.body
	presignData["noncestr"] = this.outTradeNo
	presignData["timestamp"] = strconv.Itoa(this.fee)
}

type WxPaymentNotify struct {
}

func (this *WxPaymentNotify) Notify() {

}
