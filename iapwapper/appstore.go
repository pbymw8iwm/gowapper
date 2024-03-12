package iapwapper

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/awa/go-iap/appstore"
)

type AppstoreImp struct {
}

func (c *AppstoreImp) Init() error {
	return nil
}
func (c *AppstoreImp) Auth(token string) (string, error) {

	return "", nil
}
func (c *AppstoreImp) Verify(appid string, data Receptdata, isTest *bool) (map[string]ProcessOrder, error) {
	cpBillList := make(map[string]ProcessOrder)
	client := appstore.New()
	iosReq := appstore.IAPRequest{
		ReceiptData: data.Purchase, // "your receipt data encoded by base64",
	}
	payResp := &appstore.IAPResponse{}
	ctx := context.Background()
	err := client.Verify(ctx, iosReq, payResp)
	if err != nil {
		beego.Error(err)
		return cpBillList, err
	}
	beego.Informational("appstore resp:%v", payResp)
	if payResp.Status != 0 {
		beego.Error(err)
		return cpBillList, err
	}
	if payResp.Receipt.BundleID != appid {
		beego.Error(err)
		return cpBillList, err
	}

	if payResp.Environment == "Sandbox" {
		*isTest = true
	}
	for _, app_v := range payResp.Receipt.InApp {
		cpBillList[app_v.TransactionID] = ProcessOrder{
			ProductId:  app_v.ProductID,
			BillNo:     app_v.TransactionID,         //苹果没有透传的订单ID
			//OrderId:    app_v.OriginalTransactionID, //渠道订单
			//Paytime:    "",                          //支付时间
			//Createtime: "",                          //订单时间（做校验，防止刷单）
		}
	}

	return cpBillList, nil
}
