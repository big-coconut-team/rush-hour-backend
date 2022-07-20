package main

import (
	// "net/http"

	"context"
	"controller_svc/controllers"
	"controller_svc/middlewares"
	"controller_svc/utils"
	"log"

	// "scalable-final-proj/backend/product_svc/p_controllers"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"

	// "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"github.com/gin-contrib/cors"
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

	utils.Initp_client()

	go func() {
		for e := range utils.Getp_client().Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n, order created",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET, POST"},
		AllowHeaders:     []string{"Origin, X-Requested-With, content-type, Authorization"},
		AllowCredentials: true,
	}))

	public := router.Group("/api")

	public.POST("/signup", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/api/user")

	protected.Use(middlewares.JwtAuthMiddleware())

	// protected.POST("/changepass", controllers.ChangePassword)
	// protected.GET("/profile", controllers.CurrentUser)

	protected.POST("/add_product", controllers.AddProduct)
	protected.GET("/list_product", controllers.DownloadPhoto)

	public.POST("/place_order", controllers.PlaceOrder)
	protected.POST("/make_payment", controllers.Pay)
	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	router.Run("localhost:8088")
}
