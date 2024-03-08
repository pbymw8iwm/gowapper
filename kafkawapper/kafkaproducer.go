package kafkawapper

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/astaxie/beego"
	//cluster "github.com/bsm/sarama-cluster"
)

type KafkaProducerProxy struct {
	Producer *sarama.SyncProducer
	cb       ComsumerMessageCallback
}

func (k *KafkaConsumerProxy) SetConsumeCallback(cb ComsumerMessageCallback) {
	//buxuyao
}
func (k *KafkaProducerProxy) Start(hosts []string, password string, username string, groupid string, topic []string) error {
	var err error
	config := cluster.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForLocal        // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	if password != "" || username != "" {
		// sasl认证
		config.Net.SASL.Enable = true
		config.Net.SASL.User = username
		config.Net.SASL.Password = password
	}
	k.Producer, err = sarama.NewSyncProducer(hosts, config)
	if k.Producer == nil {
		panic(fmt.Sprintf("consumer is nil. kafka info -> {brokers:%v, topic: %v, group: %v}", hosts, topic, groupid))
	}
	beego.Informational("kafka init success, consumer -> %v, topic -> %v", k.Producer, topic)
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
func (k *KafkaProducerProxy) Loop() {

}
func (k *KafkaProducerProxy) Stop() error {
	return k.Producer.Close()
}
