package main

import (
	"context"
	"log"
	"user-service/entrypoints"
	"user-service/models"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initProvider() func() {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("user"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	traceExporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint("localhost:4318"),
	)
	if err != nil {
		log.Fatal(err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func() {
		err := tracerProvider.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	tp := initProvider()
	defer tp()

	models.ConnectDatabase()

	router := gin.Default()

	router.Use(otelgin.Middleware("user"))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET, POST"},
		AllowHeaders:     []string{"Origin, X-Requested-With, content-type, Authorization"},
		AllowCredentials: true,
	}))
	router.POST("/signup", entrypoints.Register)
	router.POST("/verifypassword", entrypoints.VerifyPassword)
	router.POST("/updateuser", entrypoints.ChangePassword)
	router.POST("/getuser", entrypoints.GetUser)
	router.GET("/profile", entrypoints.CurrentUser)

	router.Run("localhost:8000")
}
