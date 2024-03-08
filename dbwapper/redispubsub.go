package dbwapper

import (
	"github.com/go-redis/redis/v8"
)

//订阅消息的接收回调函数
type SubMessageCallback func(string, string) // 定义回调函数的类型，接受一个int和一个string作为参数
//redis发布订阅处理类
type RedisPubSub struct {
	//发布订阅模式
	Pubsub            *redis.PubSub
	PubsubMessageChan <-chan *redis.Message
}

func (p *RedisPubSub) SetSubscribe(psub *redis.PubSub, cb SubMessageCallback) error {
	p.Pubsub = psub
	p.PubsubMessageChan = p.Pubsub.Channel()
	go p.ProcSubscribeMsg(cb)
	return nil
}
func (p *RedisPubSub) Close() {
	p.Pubsub.Close()
}

func (p *RedisPubSub) ProcSubscribeMsg(cb SubMessageCallback) {
	for msg := range p.PubsubMessageChan {
		cb(msg.Channel, msg.Payload)
	}
}
