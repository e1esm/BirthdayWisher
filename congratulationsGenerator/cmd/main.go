package main

import (
	"congratulationsGenerator/internal/service"
	"congratulationsGenerator/utils"
	"context"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/joho/godotenv"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
)

func main() {
	utils.InitLogger()
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error while reading from env file")
	}
	token := os.Getenv("AI_TOKEN")
	port := os.Getenv("GRPC_PORT")
	name := os.Getenv("AI_CALLER_CONTAINER_NAME")
	metricsPort := os.Getenv("METRICS_PORT")
	client := openai.NewClient(token)
	if err != nil {
		utils.Logger.Fatal("Address can't be listened", zap.String("error", err.Error()))
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor))
	aiService := service.NewOpenAIService(client)
	gen_proto.RegisterCongratulationServiceServer(grpcServer, aiService)
	grpc_prometheus.Register(grpcServer)

	group := run.Group{}

	group.Add(func() error {
		server, err := net.Listen("tcp", name+port)
		if err != nil {
			return err
		}
		return grpcServer.Serve(server)
	}, func(err error) {
		grpcServer.GracefulStop()
	})
	httpServer := &http.Server{Addr: name + metricsPort}
	group.Add(func() error {
		m := http.NewServeMux()
		m.Handle("/metrics", promhttp.Handler())
		httpServer.Handler = m
		return httpServer.ListenAndServe()
	}, func(err error) {
		if err := httpServer.Close(); err != nil {
			utils.Logger.Fatal("Error while stopping the server", zap.String("msg", "Failed to stop the web server"), zap.String("error", err.Error()))
		}
	})

	group.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))
	if err := group.Run(); err != nil {
		utils.Logger.Fatal("Error while running a group", zap.String("err", err.Error()))
	}
}
