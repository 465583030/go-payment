package payment

import (
	"strconv"
	"fmt"
	"github.com/cplusgo/go-payment/helper"
	"bytes"
	"net/http"
	"io/ioutil"
	"unsafe"
)

// 参考官方文档:
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

//已经验证过
func (this *WxPaymentSigned) Unifiedorder() {
	presignData := make(map[string]string)
	presignData["appid"] = this.appId
	presignData["attach"] = this.attach
	presignData["body"] = this.body
	presignData["mch_id"] = this.mchId
	presignData["nonce_str"] = this.nonceStr
	presignData["notify_url"] = this.notifyUrl
	presignData["out_trade_no"] = this.outTradeNo
	presignData["spbill_create_ip"] = this.createIp
	presignData["total_fee"] = strconv.Itoa(this.fee)
	presignData["trade_type"] = this.tradeType
	params := helper.ToURLParamsSortByKey(presignData)
	var buf bytes.Buffer
	buf.WriteString(params)
	buf.WriteString("&key=")
	buf.WriteString(this.appKey)
	params = buf.String()
	presignData["sign"] = helper.MD5(params)
	xml := helper.MapToXMLString(presignData)
	fmt.Println(xml)
	reader := bytes.NewReader([]byte(xml))
	resp, err := http.Post(wxUnifiedorderURL, "application/xml", reader)
	if err == nil {
		respBytes, respErr := ioutil.ReadAll(resp.Body)
		if respErr == nil {
			fmt.Println(*(*string)(unsafe.Pointer(&respBytes)))
		}
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
