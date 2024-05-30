package demo

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/dnwe/otelsarama"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var newUserCnt metric.Int64Counter

func NewUser(producer sarama.SyncProducer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		// Create a new span from cuurent context
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
		ctx, span2 := span1.TracerProvider().Tracer("interceptors").Start(r.Context(), "Send message to kafka")
		defer span2.End()

		// Create message
		id := 1 + rand.Intn(100)
		msg := sarama.ProducerMessage{
			Topic: "newuser",
			Key:   sarama.StringEncoder(fmt.Sprintf("New User with id=%d", id)),
			Value: sarama.StringEncoder(fmt.Sprintf("%d", id)),
		}
		// Inject the span context into the message
		otel.GetTextMapPropagator().Inject(ctx, otelsarama.NewProducerMessageCarrier(&msg))

		// Send message
		partition, offset, err := producer.SendMessage(&msg)
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
