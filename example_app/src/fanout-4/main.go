package main

import (
	"context"
	"log"
	"net"
	"os"
  "fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/hanapedia/microservice-topologies/tracing-example-app/fanout-4/connections/mongo"
	pb "github.com/hanapedia/microservice-topologies/tracing-example-app/fanout-4/pb-fanout-4"
)

type fanout_4Server struct {
  mongoConnection *mongo.Mongo
	pb.UnimplementedFanout_4Server
}


const (
	ListenPort = "4005"
	DbAddress  = "mongodb://localhost:27017"
  DbUser = "root"
  DbPassword = "example"
  DbName = "tracing-example-app"
  CollectionName = "fanout4"
)

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
	dbAddress := DbAddress
	if os.Getenv("DB_ADDRESS") != "" {
		dbAddress = os.Getenv("DB_ADDRESS")
	}
	dbUser := DbUser
	if os.Getenv("DB_USER") != "" {
		dbUser = os.Getenv("DB_USER")
	}
	dbPassword := DbPassword
	if os.Getenv("DB_PASSWORD") != "" {
		dbPassword = os.Getenv("DB_PASSWORD")
	}
	dbName := DbName
	if os.Getenv("DB_NAME") != "" {
		dbName = os.Getenv("DB_NAME")
	}
	collectionName := CollectionName
	if os.Getenv("COLLECTION_NAME") != "" {
		collectionName = os.Getenv("COLLECTION_NAME")
	}

  // ADDED
	tp := InitTracerProvider()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	mongoConnection, err := mongo.InitMongo(dbAddress, dbUser, dbPassword, dbName, collectionName)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer mongoConnection.Disconnect()

	svc := new(fanout_4Server)
  svc.mongoConnection = mongoConnection
  lis, err := net.Listen("tcp", fmt.Sprintf(":%s", listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
  opts = append(opts, 
    grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
    grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
    )
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFanout_4Server(grpcServer, svc)
  healthpb.RegisterHealthServer(grpcServer, svc)

	grpcServer.Serve(lis)
}

func (s fanout_4Server) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
	newId, err := s.mongoConnection.GetItem(req.Ids[len(req.Ids)-1])
	if err != nil {
		log.Printf("Failed to retrieve item at fanout_4: %v", err)
	}
	returnArr := req.Ids
	returnArr = append(returnArr, newId)
	// Implement db read logic
	res := pb.Res{
		Ids: returnArr,
	}

	return &res, nil
}

func (s fanout_4Server) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (s fanout_4Server) Watch(req *healthpb.HealthCheckRequest, ws healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
