package kafkawapper

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/astaxie/beego"
	//cluster "github.com/bsm/sarama-cluster"
)

type KafkaProducerParam struct {
	User         string
	Password     string
	Server_addrs []string
}

//参数分别是  mastername， password， dbindex ，server_addrs
func (p *KafkaProducerParam) GetCfg() interface{} {
	return p
}

type KafkaProducerProxy struct {
	Producer sarama.SyncProducer
}

func (k *KafkaProducerProxy) Close() error {
	return k.Producer.Close()
}
func (k *KafkaProducerProxy) Start(param *KafkaProducerParam) error {
	var err error //
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForLocal        // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	if param.Password != "" || param.User != "" {
		// sasl认证
		config.Net.SASL.Enable = true
		config.Net.SASL.User = param.User
		config.Net.SASL.Password = param.Password
	}
	k.Producer, err = sarama.NewSyncProducer(param.Server_addrs, config)
	if k.Producer == nil {
		panic(fmt.Sprintf("consumer is nil. kafka info -> {brokers:%v}", param.Server_addrs))
	}
	beego.Informational("kafka init success, err:%v consumer -> %v", err, k.Producer)
	return err
}

func (k *KafkaProducerProxy) ProduceMsg(topic string, msg string) error {
	msgX := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	fmt.Printf("SendMsg -> topic:%v-%v\n", topic, msg)
	if k.Producer != nil {
		// 发送消息
		partition, offset, err := k.Producer.SendMessage(msgX)
		if err != nil {
			fmt.Printf("send msg error:%s \n", partition, offset, err)
		}
		return err
	}
	return nil

}
func (k *KafkaProducerProxy) Stop() error {
	return k.Producer.Close()
}
