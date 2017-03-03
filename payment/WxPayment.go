package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cplusgo/go-payment/helper"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"unsafe"
	"encoding/xml"
)

// 参考官方文档:
// 统一下单接口:https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1
// 调起支付:https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12&index=2

const wxUnifiedorderURL string = "https://api.mch.weixin.qq.com/pay/unifiedorder"

/**
<xml>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str>
   <sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id>
   <trade_type><![CDATA[APP]]></trade_type>
</xml>
*/

type UnifiedorderResult struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	PrepayId   string `xml:"prepay_id"`
	TradeType  string `xml:"trade_type"`
}

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
	xmlString := helper.MapToXMLString(presignData)
	fmt.Println(xmlString)
	reader := bytes.NewReader([]byte(xmlString))
	resp, err := http.Post(wxUnifiedorderURL, "application/xml", reader)
	var respXmlString string
	var result UnifiedorderResult
	if err == nil {
		respBytes, respErr := ioutil.ReadAll(resp.Body)
		if respErr == nil {
			respXmlString = *(*string)(unsafe.Pointer(&respBytes))
			fmt.Println(respXmlString)
			xml.Unmarshal(respBytes, &result)
		}
	}
	fmt.Println(result.PrepayId)

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
