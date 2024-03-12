package iapwapper

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"loginserver/protocol"
	"time"

	"github.com/astaxie/beego"
	"github.com/awa/go-iap/playstore"
	"google.golang.org/api/idtoken"
)

type GoogleplayConfigModal struct {
	Web GoogleplayConfigWeb `json:"web"`
}

type GoogleplayConfigWeb struct {
	AuthProviderX509CERTURL string `json:"auth_provider_x509_cert_url,omitempty"`
	AuthURI                 string `json:"auth_uri"`
	ClientID                string `json:"client_id"`
	ClientSecret            string `json:"client_secret"`
	ProjectID               string `json:"project_id"`
	TokenURI                string `json:"token_uri"`
}

type GoogleplayImp struct {
	data    GoogleplayConfigModal
	config  []byte
	iapjson []byte
}

func (c *GoogleplayImp) Init() error {
	config_path := beego.AppConfig.String("googleplay::config_path")
	config, err := ioutil.ReadFile(config_path)
	if err != nil {
		beego.Error(err)
		return err
	}
	c.config = config
	//beego.Informational("read google config ", string(c.config))
	err = json.Unmarshal(c.config, &c.data)
	if err != nil {
		beego.Error(err)
		return err
	}
	iap_path := beego.AppConfig.String("googleplay::iap_path")
	c.iapjson, err = ioutil.ReadFile(iap_path)
	if err != nil {
		beego.Error(err)
		return err
	}
	return nil
}
func (c *GoogleplayImp) Auth(token string) (string, error) {
	payload, err := idtoken.Validate(context.Background(), token, c.data.Web.ClientID)
	beego.Informational("get google login token verify resp %v, error=%v", payload, err)
	if err != nil {
		return "", err
	}

	return payload.Claims["sub"].(string), nil
}
func (c *GoogleplayImp) Verify(appid string, data protocol.Receptdata, isTest *bool) (map[string]protocol.ProcessOrder, error) {
	cpBillList := make(map[string]protocol.ProcessOrder)

	client, Err := playstore.New(c.iapjson)
	if Err != nil {
		beego.Error(Err)
		return cpBillList, Err
	}
	ctx := context.Background()
	i := 0
	for {
		i = i + 1
		if data.PayType == 0 {
			payResp, err := client.VerifyProduct(ctx, appid, data.Productid, data.Token)
			beego.Informational("get google token verify resp %v, error=%v", payResp, err)
			if err != nil {
				beego.Error(err)
				return cpBillList, err
			}

			beego.Informational("playstore resp:%v", payResp.ProductId, ",PurchaseState:%v", payResp.PurchaseState, ",DeveloperPayload:", payResp.DeveloperPayload)
			bs, _ := json.Marshal(*payResp)
			beego.Informational("playstore resp:%v", string(bs))
			fmt.Println("test: %+v", payResp)

			// PurchaseState: The purchase state of the order. Possible values are:
			// 0. Purchased 1. Canceled 2. Pending
			if payResp.PurchaseState == 2 {
				if i > 20 {
					break
				}
				time.Sleep(1000 * time.Millisecond)
				continue
			}
			if payResp.PurchaseState == 0 {
				cpBillList[payResp.OrderId] = protocol.ProcessOrder{
					ProductId: data.Productid,
					BillNo:    payResp.ObfuscatedExternalAccountId, //gpOrder.DeveloperPayload  开发者自定义上传的订单号
				}
			}
			// PurchaseType: The type of purchase of the inapp product. This field
			// is only set if this purchase was not made using the standard in-app
			// billing flow. Possible values are: 0. Test (i.e. purchased from a
			// license testing account) 1. Promo (i.e. purchased using a promo code)
			// 2. Rewarded (i.e. from watching a video ad instead of paying)
			if *payResp.PurchaseType == 0 {
				*isTest = true
			}
		} else {
			payResp, err := client.VerifySubscription(ctx, appid, data.Productid, data.Token)
			beego.Informational("get google token verify resp %v, error=%v", payResp, err)
			if err != nil {
				beego.Error(err)
				return cpBillList, err
			}

			bs, _ := json.Marshal(payResp)
			beego.Informational("playstore resp:%v", string(bs))
			//PaymentState: The payment state of the subscription. Possible values
			//are: 0. Payment pending 1. Payment received 2. Free trial 3. Pending
			//deferred upgrade/downgrade Not present for canceled, expired
			//subscriptions.
			if *payResp.PaymentState == 0 || *payResp.PaymentState == 3 {
				if i > 20 {
					break
				}
				time.Sleep(1000 * time.Millisecond)
				continue
			}
			if *payResp.PaymentState == 1 || *payResp.PaymentState == 2 {
				cpBillList[payResp.OrderId] = protocol.ProcessOrder{
					ProductId: data.Productid,
					BillNo:    payResp.ObfuscatedExternalAccountId, //payResp.DeveloperPayload,  //gpOrder.DeveloperPayload  开发者自定义上传的订单号
				}
			}
			if *payResp.PurchaseType == 0 {
				*isTest = true
			}

		}
		break
	}
	return cpBillList, nil
}
