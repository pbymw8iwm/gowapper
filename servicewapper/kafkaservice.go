package servicewapper

import (
	"github.com/pbymw8iwm/gowapper/kafkawapper"
)

type KafkaManager struct {
	Producer *kafkawapper.KafkaProducerProxy
}

func (p *KafkaManager) Stop() error {
	err := p.Producer.Close()
	p.Producer = nil
	return err
}

func (p *KafkaManager) Start(params IServiceParams) error {
	p.Producer = new(kafkawapper.KafkaProducerProxy)
	err := p.Producer.Start((params.GetCfg().(*kafkawapper.KafkaProducerParam)))
	if err != nil {
		return err
	}
	return nil
}
func (p *KafkaManager) SendMsg(topic, msg string) error {
	err := p.Producer.ProduceMsg(topic, msg)
	if err != nil {
		return err
	}
	return nil
}
func test_kafka() {
	gkafka := &(KafkaManager{})
	//参数分别是  mastername， password， dbindex ，server_addrs  192.168.10.214:6380
	var param IServiceParams = &kafkawapper.KafkaProducerParam{
		User: "mymaster", Password: "", Server_addrs: []string{"192.168.10.214:6380"},
	}
	gkafka.Start(param)
	gkafka.SendMsg("test", "saassss")
}
