package demo

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"strconv"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

const name = "hello-service"

var (
	tracer  = otel.Tracer(name)
	meter   = otel.Meter(name)
	logger  = otelslog.NewLogger(name)
	rollCnt metric.Int64Counter
)

func GetRandomData(w http.ResponseWriter, r *http.Request) {
	var err error
	number := 1 + rand.Intn(6)

	// Trace
	_, span := tracer.Start(r.Context(), "GetRandomData")
	defer span.End()

	// Metric
	rollCnt, _ = meter.Int64Counter("random.counter",
		metric.WithDescription("The number of random numbers"),
		metric.WithUnit("{roll}"))

	rollCnt.Add(r.Context(), 1)

	// Log
	logger.Info(fmt.Sprintf("Rolled a %d\n", number), slog.Attr{Key: "traceid", Value: slog.StringValue(span.SpanContext().TraceID().String())})

	resp := strconv.Itoa(number)
	if _, err = io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
