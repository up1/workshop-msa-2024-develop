#!/bin/bash

set -euo pipefail
export OTEL_METRIC_EXPORT_INTERVAL="5000"  # so we don't have to wait 60s for metrics
export OTEL_RESOURCE_ATTRIBUTES="service.name=hello,service.instance.id=localhost:9000"

# Run the application
export OTEL_EXPORTER_OTLP_INSECURE="true" # use http instead of https (needed because of https://github.com/open-telemetry/opentelemetry-go/issues/4834)
go run cmd/main.go