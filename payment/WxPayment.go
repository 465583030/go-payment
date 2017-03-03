package payment

import (
	"strconv"
	"fmt"
	"github.com/cplusgo/go-payment/helper"
	"bytes"
	"net/http"
	"io/ioutil"
	"unsafe"
	"time"
	"encoding/json"
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

//已经验证过,这一步的作用是拿到预支付交易会话ID,即prepayid
func (this *WxPaymentSigned) Unifiedorder() string {
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
	var respXmlString string
	if err == nil {
		respBytes, respErr := ioutil.ReadAll(resp.Body)
		if respErr == nil {
			respXmlString = *(*string)(unsafe.Pointer(&respBytes))
			fmt.Println(respXmlString)
		}
	}
	return respXmlString
}

//在统一下单接口执行完毕之后会返回一个关键的数据,prepayid
func (this *WxPaymentSigned) Signed(prepayid string) []byte {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	presignData := make(map[string]string)
	presignData["appid"] = this.appId
	presignData["partnerid"] = this.mchId
	presignData["prepayid"] = prepayid
	presignData["package"] = "Sign=WXPay"
	presignData["noncestr"] = helper.MD5(timestamp)
	presignData["timestamp"] = timestamp

	params := helper.ToURLParamsSortByKey(presignData)
	var buf bytes.Buffer
	buf.WriteString(params)
	buf.WriteString("&key=")
	buf.WriteString(this.appKey)
	presignData["sign"] = helper.MD5(buf.String())
	jsonBody, _ := json.Marshal(presignData)
	return jsonBody
}

type WxPaymentNotify struct {
}

func (this *WxPaymentNotify) Notify() {

}
