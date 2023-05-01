package main

import (
	"context"
	"github.com/e1esm/protobuf/bridge_to_PDF-Generator/gen_proto"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/joho/godotenv"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"pdfGenerator/internal/service"
	"pdfGenerator/utils"
	"syscall"
)

func main() {
	utils.InitLogger()
	err := godotenv.Load("../.env")
	if err != nil {
		utils.Logger.Fatal("Couldn't have opened env file", zap.String("err", err.Error()))
	}
	grpcPort := os.Getenv("GRPC_PORT")
	address := os.Getenv("pdf_generator_container_name")
	metricsPort := os.Getenv("METRICS_PORT")

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor))
	gen_proto.RegisterPDFGenerationServiceServer(grpcServer, &service.PDFService{})
	group := run.Group{}

	group.Add(func() error {
		server, err := net.Listen("tcp", address+grpcPort)
		if err != nil {
			return err
		}
		return grpcServer.Serve(server)
	}, func(err error) {
		grpcServer.GracefulStop()
	})
	httpServer := &http.Server{Addr: address + metricsPort}
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
