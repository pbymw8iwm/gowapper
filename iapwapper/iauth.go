package iapwapper

import (
	"loginserver/protocol"
)

type Receptdata struct {
	Purchase  string `json:"purchase,omitempty"`
	Token     string `json:"token,omitempty"`
	Type      string `json:"type"`
	Productid string `json:"productid,omitempty"` //目前不用传
	Billno    string `json:"billno,omitempty"`    //目前不用传
	PayType   int32  `json:"paytype"`
}
type ProcessOrder struct {
	ProductId  string //产品iD
	BillNo     string //游戏服务器的订单号
	OrderId    string //渠道订单
	Paytime    string //支付时间
	Createtime string //订单时间（做校验，防止刷单）
}

type AuthInterface interface {
	Init() error
	Auth(token string) (string, error)
}

//* AppStore: [![GoDoc](https://godoc.org/github.com/awa/go-iap/appstore?status.svg)](https://godoc.org/github.com/awa/go-iap/appstore)
//* GooglePlay: [![GoDoc](https://godoc.org/github.com/awa/go-iap/playstore?status.svg)](https://godoc.org/github.com/awa/go-iap/playstore)
//* Amazon AppStore: [![GoDoc](https://godoc.org/github.com/awa/go-iap/amazon?status.svg)](https://godoc.org/github.com/awa/go-iap/amazon)
//* Huawei HMS: [![GoDoc](https://godoc.org/github.com/awa/go-iap/hms?status.svg)](https://godoc.org/github.com/awa/go-iap/hms)
type IapInterface interface {
	Init() error
	Verify(appid string, data protocol.Receptdata, isTest *bool) (map[string]protocol.ProcessOrder, error)
}
