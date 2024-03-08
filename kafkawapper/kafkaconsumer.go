package kafkawapper

/*
import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/astaxie/beego"
	//cluster "github.com/bsm/sarama-cluster"
)

type KafkaConsumerProxy struct {
	Consumer *cluster.Consumer
	cb       ComsumerMessageCallback
}

func (k *KafkaConsumerProxy) SetConsumeCallback(cb ComsumerMessageCallback) {
	k.cb = cb
}

// deserializeMessage 函数用于反序列化消息
func deserializeMessage(data []byte) (*Message, error) {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
func (k *KafkaConsumerProxy) Start(hosts []string, password string, username string, groupid string, topic []string) error {
	var err error
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = -2
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Group.Return.Notifications = true
	if password != "" || username != "" {
		// sasl认证
		config.Net.SASL.Enable = true
		config.Net.SASL.User = username
		config.Net.SASL.Password = password
	}
	k.Consumer, err = sarama.NewConsumer(hosts, config)
	if err != nil {
		panic(err.Error())
	}
	if k.Consumer == nil {
		panic(fmt.Sprintf("consumer is nil. kafka info -> {brokers:%v, topic: %v, group: %v}", hosts, topic, groupid))
	}
	// 订阅主题
	partition := int32(0)
	offset := int64(0)

	// 创建分区消费者
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %s", err.Error())
	}

	// 开启消费者组协程
	go func() {
		// 消费消息
		for message := range partitionConsumer.Messages() {
			// 反序列化消息
			deserializedMessage, err := deserializeMessage(message.Value)
			if err != nil {
				log.Printf("Error deserializing message: %s", err.Error())
				continue
			}

			// 在这里处理反序列化后的消息
			fmt.Printf("Received message: ID=%d, Data=%s\n", deserializedMessage.ID, deserializedMessage.Data)
		}
	}()
	beego.Informational("kafka init success, consumer -> %v, topic -> %v", k.Consumer, topic)
}
func (k *KafkaConsumerProxy) ProduceMsg(topic string, msg string) error {
	return nil
}



func (k *KafkaConsumerProxy) Stop() error {
	return k.Consumer.Close()
}*/
