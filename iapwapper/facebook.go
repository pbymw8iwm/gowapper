package iapwapper

import (
	"loginserver/protocol"

	"github.com/astaxie/beego"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type FacebookAccessModal struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type FacebookAuthModal struct {
	FacebookAuthData FacebookAuthData `json:"data"`
}

type FacebookAuthData struct {
	AppID               string   `json:"app_id"`
	Application         string   `json:"application"`
	DataAccessExpiresAt int64    `json:"data_access_expires_at"`
	ExpiresAt           int64    `json:"expires_at"`
	IsValid             bool     `json:"is_valid"`
	IssuedAt            int64    `json:"issued_at"`
	Metadata            Metadata `json:"metadata"`
	Scopes              []string `json:"scopes"`
	Type                string   `json:"type"`
	UserID              string   `json:"user_id"`
}

type Metadata struct {
	AuthType string `json:"auth_type"`
	Sso      string `json:"sso"`
}

type FacebookImp struct {
	client_id     string
	client_secret string
	access_url    string
	token_url     string
}

func (c *FacebookImp) Init() error {
	c.client_id = beego.AppConfig.String("facebook::client_id")
	c.client_secret = beego.AppConfig.String("facebook::client_secret")
	c.access_url = "https://graph.facebook.com/oauth/access_token"
	c.token_url = "https://graph.facebook.com/debug_token"
	return nil
}

func (c *FacebookImp) GetAccessToken() (string, error) {
	method := "GET"
	url := fmt.Sprintf("%s?client_id=%s&client_secret=%s&grant_type=client_credentials", c.access_url, c.client_id, c.client_secret)
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*5) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 60)) //设置发送接受数据超时
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 5, //设置读取头部超时client.Do(req)
			//DisableKeepAlives: true,//为TRUE的话每次client.Do(req)都会建立新的连接，FALSE的话会保持长连接（使用连接池），默认为FALSE
		},
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		beego.Informational(err)
		return "", err
	}
	req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")

	res, err := client.Do(req)
	if err != nil {
		beego.Informational(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		beego.Informational(err)
		return "", err
	}
	beego.Informational("get google token %v, error=%v", string(body), err)
	info := FacebookAccessModal{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return "", err
	}
	return info.AccessToken, nil
}
func (c *FacebookImp) VerifyInputToken(input_token string, access_token string) (string, error) {

	url := fmt.Sprintf("%s?input_token=%s&access_token=%s", c.token_url, input_token, access_token)
	method := "GET"

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*5) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 60)) //设置发送接受数据超时
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 5, //设置读取头部超时client.Do(req)
			//DisableKeepAlives: true,//为TRUE的话每次client.Do(req)都会建立新的连接，FALSE的话会保持长连接（使用连接池），默认为FALSE
		},
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		beego.Informational(err)
		return "", err
	}
	req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")

	res, err := client.Do(req)
	if err != nil {
		beego.Informational(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		beego.Informational(err)
		return "", err
	}
	beego.Informational("get google token %v, error=%v", string(body), err)
	info := FacebookAuthModal{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return "", err
	}
	return info.FacebookAuthData.UserID, nil
}
func (c *FacebookImp) Auth(token string) (string, error) {
	access_token, err := c.GetAccessToken()
	if err != nil {
		return "", err
	}
	//access_token := fmt.Sprintf("%s|%s",c.client_id, c.client_secret)
	user_id, err := c.VerifyInputToken(token, access_token)
	return user_id, err
}
func (c *FacebookImp) Verify(appid string, data protocol.Receptdata, isTest *bool) (map[string]protocol.ProcessOrder, error) {
	cpBillList := make(map[string]protocol.ProcessOrder)
	return cpBillList, nil
}
