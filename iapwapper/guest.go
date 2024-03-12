package iapwapper

import (
	"crypto/md5"
	"fmt"
	"loginserver/protocol"

	"github.com/astaxie/beego"
)

type GuestUserImp struct {
	prefix string
	md5mix string
}

func (c *GuestUserImp) Init() error {
	c.prefix = "tmhy"
	c.md5mix = "hello world, tmhy no1."
	return nil
}
func (c *GuestUserImp) Auth(token string) (string, error) {
	beego.Informational("guest auth token:%s", token)
	if token == "" {
		return "", fmt.Errorf("invalid guest token")
	}
	data := []byte(fmt.Sprintf("%s%s", token, c.md5mix))
	has := md5.Sum(data)
	user_id := fmt.Sprintf("%s-%x", c.prefix, has)
	beego.Informational("guest auth userid:%s", user_id)
	return user_id, nil
}
func (c *GuestUserImp) Verify(appid string, data protocol.Receptdata, isTest *bool) (map[string]protocol.ProcessOrder, error) {
	cpBillList := make(map[string]protocol.ProcessOrder)

	return cpBillList, nil
}
