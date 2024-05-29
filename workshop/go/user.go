package demo

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"net/http"

	"github.com/IBM/sarama"
	"go.opentelemetry.io/otel/metric"
)

var newUserCnt metric.Int64Counter

func NewUser(producer sarama.SyncProducer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		// Trace
		_, span1 := tracer.Start(r.Context(), "Call NewUser")
		defer span1.End()

		// Metric
		newUserCnt, _ = meter.Int64Counter("newuser.counter",
			metric.WithDescription("The number of new user"),
			metric.WithUnit("{person}"))

		newUserCnt.Add(r.Context(), 1)

		// Log
		logger.Info("Publish data to kafka", slog.Attr{Key: "traceid", Value: slog.StringValue(span1.SpanContext().TraceID().String())})

		// Create new span from the parent span
		_, span2 := span1.TracerProvider().Tracer("interceptors").Start(r.Context(), "Send message to kafka")
		defer span2.End()

		id := 1 + rand.Intn(100)
		partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
			Topic: "newuser",
			Value: sarama.StringEncoder(fmt.Sprintf("New User with id=%d", id)),
		})
		if err != nil {
			log.Panicf("Error from consumer: %v", err)
		} else {
			log.Printf("Your data is stored with unique identifier quickstart/%d/%d\n", partition, offset)
		}

		resp := "New user created"
		if _, err = io.WriteString(w, resp); err != nil {
			log.Printf("Write failed: %v\n", err)
		}
	}
}
