package tracer

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func InitProvider() func() {
	exporter, err := zipkin.New(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"))
	if err != nil {
		log.Fatalf("failed to create Zipkin exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(os.Getenv("OTEL_SERVICE_NAME")),
		)),
	)
	otel.SetTracerProvider(tp)
	return func() {
		_ = tp.Shutdown(context.Background())
	}
}

func Tracer() trace.Tracer {
	return otel.Tracer(os.Getenv("OTEL_SERVICE_NAME"))
}
