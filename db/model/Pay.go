package model

type PayInfo struct {
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
	Respdata PayDetail `json:"respdata"`
}

type PayDetail struct {
	ReturnCode string `json:"return_code"`
	ReturnMsg string `json:"return_msg"`
	Appid     string `json:"appid"`
	MchId  string `json:"mch_id"`
	SubAppid string `json:"sub_appid"`
	SubMchId string `json:"sub_mch_id"`
	NonceStr string `json:"nonce_str"`
	Sign string `json:"sign"`
	ResultCode string `json:"result_code"`
	TradeType string `json:"trade_type"`
	PrepayId string `json:"prepay_id"`
	Payment PaymentInfo `json:"payment"`
}

type PaymentInfo struct {
	AppId string `json:"app_id"`
	TimeStamp string `json:"timeStamp"`
	NonceStr string `json:"nonceStr"`
	Package string `json:"package"`
	SignType string `json:"signType"`
	PaySign string `json:"paySign"`
}
