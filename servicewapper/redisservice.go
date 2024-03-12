package servicewapper

import (
	"github.com/astaxie/beego"
	"github.com/pbymw8iwm/gowapper/dbwapper"
)

func PrintSubscribeMsg(channel, payload string) {
	beego.Informational("redis pubsub recv ->", channel, payload)
}

type RedisManager struct {
	RedisClient dbwapper.IRedisClient
	//发布订阅模式
	RedisPubSub *dbwapper.RedisPubSub
}

func (p *RedisManager) Stop() error {
	p.RedisPubSub.Close()
	err := p.RedisClient.Close()
	p.RedisClient = nil
	if p.RedisPubSub != nil {
		p.RedisPubSub = nil
	}

	return err
}

func (p *RedisManager) Start(params IServiceParams) error {
	p.RedisClient = new(dbwapper.RedisClusterClient)
	err := p.RedisClient.Connect((params.GetCfg().(*dbwapper.RedisCacheParam)))
	if err != nil {
		return err
	}
	//p.RedisPubSub = new(dbwapper.RedisPubSub)
	//p.RedisPubSub.SetSubscribe(p.RedisClient.Subscribe("test"), PrintSubscribeMsg)
	return nil
}

func test_redis() {
	gredis := &(RedisManager{})
	//参数分别是  mastername， password， dbindex ，server_addrs  192.168.10.214:6380
	var redispaa IServiceParams = &dbwapper.RedisCacheParam{
		Mastername: "mymaster", Password: "", Dbindex: 1, Server_addrs: []string{"192.168.10.214:6380"},
	}
	gredis.Start(redispaa)

	gredis.RedisClient.Publish("test", "hello321")
}
