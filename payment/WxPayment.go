package payment

type WxPaymentSigned struct {
	appId     string
	appKey    string
	mchId     string
	nonceStr  string
	desc      string
	fee       int64
	notifyUrl string
	orderId   int64
}

func (this *WxPaymentSigned) Signed() {

}

type WxPaymentNotify struct {
}

func (this *WxPaymentNotify) Notify() {

}
