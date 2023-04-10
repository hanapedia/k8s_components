package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
  "math/rand"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

  "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	pb1 "github.com/hanapedia/microservice-topologies/tracing-example-app/gateway/pb-fanout-1"
	pb2 "github.com/hanapedia/microservice-topologies/tracing-example-app/gateway/pb-fanout-2"
	pb3 "github.com/hanapedia/microservice-topologies/tracing-example-app/gateway/pb-fanout-3"
	pb4 "github.com/hanapedia/microservice-topologies/tracing-example-app/gateway/pb-fanout-4"
	pb5 "github.com/hanapedia/microservice-topologies/tracing-example-app/gateway/pb-fanout-5"
)

const (
	ListenPort           = "4000"
  ServerName = "gateway"
	FanoutClient1Address = "localhost:4002"
	FanoutClient2Address = "localhost:4003"
	FanoutClient3Address = "localhost:4004"
	FanoutClient4Address = "localhost:4005"
	FanoutClient5Address = "localhost:4006"
)

type gatewayServer struct {
  fanout_1Client pb1.Fanout_1Client
  fanout_2Client pb2.Fanout_2Client
  fanout_3Client pb3.Fanout_3Client
  fanout_4Client pb4.Fanout_4Client
  fanout_5Client pb5.Fanout_5Client
}

func InitTracerProvider() *sdktrace.TracerProvider {
	ctx := context.Background()

	exporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

  // must set
  // OTEL_EXPORTER_OTLP_TRACES_ENDPOINT
  // OTEL_RESOURCE_ATTRIBUTES
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func main() {
	listenPort := ListenPort
	if os.Getenv("LISTEN_PORT") != "" {
		listenPort = os.Getenv("LISTEN_PORT")
	}
	fanoutClient1Address := FanoutClient1Address
	if os.Getenv("FANOUT_CLIENT_1_ADDRESS") != "" {
		fanoutClient1Address = os.Getenv("FANOUT_CLIENT_1_ADDRESS")
	}
	fanoutClient2Address := FanoutClient2Address
	if os.Getenv("FANOUT_CLIENT_2_ADDRESS") != "" {
		fanoutClient2Address = os.Getenv("FANOUT_CLIENT_2_ADDRESS")
	}
	fanoutClient3Address := FanoutClient3Address
	if os.Getenv("FANOUT_CLIENT_3_ADDRESS") != "" {
		fanoutClient3Address = os.Getenv("FANOUT_CLIENT_3_ADDRESS")
	}
	fanoutClient4Address := FanoutClient4Address
	if os.Getenv("FANOUT_CLIENT_4_ADDRESS") != "" {
		fanoutClient4Address = os.Getenv("FANOUT_CLIENT_4_ADDRESS")
	}
	fanoutClient5Address := FanoutClient5Address
	if os.Getenv("FANOUT_CLIENT_5_ADDRESS") != "" {
		fanoutClient5Address = os.Getenv("FANOUT_CLIENT_5_ADDRESS")
	}
  serverName := ServerName
	if os.Getenv("SERVER_NAME") != "" {
		serverName = os.Getenv("SERVER_NAME")
	}

	tp := InitTracerProvider()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	var optsClient []grpc.DialOption
  optsClient = append(optsClient, 
    grpc.WithTransportCredentials(insecure.NewCredentials()),
    grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
    grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
    )

  svc := new(gatewayServer)

	conn_1, err := grpc.DialContext(context.Background(),fanoutClient1Address, optsClient...)
	if err != nil {
		log.Fatalf("Cannot establish connection with the server: %v", err)
	}
	log.Printf("Dialed and established connection with %v", fanoutClient1Address)
	defer conn_1.Close()

	svc.fanout_1Client = pb1.NewFanout_1Client(conn_1)
  log.Printf("established connection at %s", fanoutClient1Address)

	conn_2, err := grpc.DialContext(context.Background(),fanoutClient2Address, optsClient...)
	if err != nil {
		log.Fatalf("Cannot establish connection with the server: %v", err)
	}
	log.Printf("Dialed and established connection with %v", fanoutClient2Address)
	defer conn_2.Close()

	svc.fanout_2Client = pb2.NewFanout_2Client(conn_2)
  log.Printf("established connection at %s", fanoutClient2Address)

	conn_3, err := grpc.DialContext(context.Background(),fanoutClient3Address, optsClient...)
	if err != nil {
		log.Fatalf("Cannot establish connection with the server: %v", err)
	}
	log.Printf("Dialed and established connection with %v", fanoutClient3Address)
	defer conn_3.Close()

	svc.fanout_3Client = pb3.NewFanout_3Client(conn_3)
  log.Printf("established connection at %s", fanoutClient3Address)

	conn_4, err := grpc.DialContext(context.Background(),fanoutClient4Address, optsClient...)
	if err != nil {
		log.Fatalf("Cannot establish connection with the server: %v", err)
	}
	log.Printf("Dialed and established connection with %v", fanoutClient4Address)
	defer conn_4.Close()

	svc.fanout_4Client = pb4.NewFanout_4Client(conn_4)
  log.Printf("established connection at %s", fanoutClient4Address)

	conn_5, err := grpc.DialContext(context.Background(),fanoutClient5Address, optsClient...)
	if err != nil {
		log.Fatalf("Cannot establish connection with the server: %v", err)
	}
	log.Printf("Dialed and established connection with %v", fanoutClient5Address)
	defer conn_5.Close()

	svc.fanout_5Client = pb5.NewFanout_5Client(conn_5)
  log.Printf("established connection at %s", fanoutClient5Address)

  router := gin.Default()
  router.Use(otelgin.Middleware(serverName))
  router.GET("/ids", svc.handler)
  router.Run(fmt.Sprintf(":%s", listenPort))
}

func (s gatewayServer) handler(c *gin.Context) {
  res := []int32{rand.Int31n(50)}

  fanout_1Req := new(pb1.Req)
  fanout_1Req.Ids = res
	fanout_1Res, err := s.fanout_1Client.GetIds(c.Request.Context(), fanout_1Req)
	if err == nil {
		res = fanout_1Res.Ids
	} else {
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "failed",
      "message": err.Error(),
    })
    return
	}
  fanout_2Req := new(pb2.Req)
  fanout_2Req.Ids = res
	fanout_2Res, err := s.fanout_2Client.GetIds(c.Request.Context(), fanout_2Req)
	if err == nil {
		res = fanout_2Res.Ids
	} else {
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "failed",
      "message": err.Error(),
    })
    return
	}
  fanout_3Req := new(pb3.Req)
  fanout_3Req.Ids = res
	fanout_3Res, err := s.fanout_3Client.GetIds(c.Request.Context(), fanout_3Req)
	if err == nil {
		res = fanout_3Res.Ids
	} else {
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "failed",
      "message": err.Error(),
    })
    return
	}
  fanout_4Req := new(pb4.Req)
  fanout_4Req.Ids = res
	fanout_4Res, err := s.fanout_4Client.GetIds(c.Request.Context(), fanout_4Req)
	if err == nil {
		res = fanout_4Res.Ids
	} else {
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "failed",
      "message": err.Error(),
    })
    return
	}
  fanout_5Req := new(pb5.Req)
  fanout_5Req.Ids = res
	fanout_5Res, err := s.fanout_5Client.GetIds(c.Request.Context(), fanout_5Req)
	if err == nil {
		res = fanout_5Res.Ids
	} else {
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "failed",
      "message": err.Error(),
    })
    return
	}

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "message": res,
  })
}
