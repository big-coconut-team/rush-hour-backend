package main

import (
	"context"
	"log"
	"product_svc/p_controllers"
	"product_svc/p_models"

	"product_svc/p_utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	p_utils.Initp_client()

	go func() {
		for e := range p_utils.Getp_client().Events() {
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

	go p_controllers.RunQueue()

	p_models.ConnectDataBase()

	router := gin.Default()
	router.Use(otelgin.Middleware("product"))
	protected := router.Group("/api/user")
	protected.POST("/add_product", p_controllers.AddProduct)
	protected.GET("/list_product", p_controllers.DownloadPhoto)
	// protected.POST("/update_stock", p_controllers.GetStockUpdate)

	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	router.Run("localhost:8001")
}

// package main

// import (
// 	"product_svc/p_controllers"
// 	"product_svc/p_models"

// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	p_models.ConnectDataBase()

// 	router := gin.Default()
// 	protected := router.Group("/api/user")
// 	protected.POST("/add_product", p_controllers.AddProduct)
// 	protected.POST("/list_product", p_controllers.DownloadPhoto)
// 	protected.POST("/update_stock", p_controllers.GetStockUpdate)
// 	protected.POST("/update_many_stock", p_controllers.GetUpdateManyStock)

// 	// router.GET("/albums", getAlbums)
// 	// router.GET("/albums/:id", getAlbumByID)
// 	// router.POST("/albums", postAlbums)

// 	router.Run("localhost:8001")
// }
