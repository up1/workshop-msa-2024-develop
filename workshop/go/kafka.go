package demo

import (
	"log"

	"github.com/IBM/sarama"
)

func NewProducer(brokerList []string, version sarama.KafkaVersion) sarama.SyncProducer {
	// For the data collector, we are looking for strong consistency semantics.
	// Because we don't change the flush settings, sarama will try to produce messages
	// as fast as possible to keep latency low.
	config := sarama.NewConfig()
	config.Version = version
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true
	config.Net.TLS.Enable = false

	// Use OpenTelemetry instrumentation.
	config.Producer.Interceptors = []sarama.ProducerInterceptor{NewOTelInterceptor(brokerList)}

	// On the broker side, you may want to change the following settings to get
	// stronger consistency guarantees:
	// - For your broker, set `unclean.leader.election.enable` to false
	// - For the topic, you could increase `min.insync.replicas`.

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	return producer
}
