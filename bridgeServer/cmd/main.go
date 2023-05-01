package main

import (
	"bridgeServer/config"
	"bridgeServer/internal/model"
	"bridgeServer/internal/repository"
	"bridgeServer/internal/service"
	"bridgeServer/utils"
	"context"
	"fmt"
	bot_to_server_proto "github.com/e1esm/protobuf/bot_to_server/gen_proto"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/joho/godotenv"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"
)

func main() {
	utils.InitLogger()
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Couldn't have opened env file: %v", err)
	}

	cfg := config.NewConfig(dbConfiguration(), GRPCClientConfiguration())
	defer cfg.DB.DB()
	userRepository := repository.NewUserRepository(cfg.DB)
	userService := service.NewUserService(userRepository)
	gptService := service.NewGPTService(cfg.Client)
	serverImpl := config.NewServer(userService, gptService, cfg)

	port := os.Getenv("GRPC_PORT")
	address := os.Getenv("BRIDGE_SERVER_CONTAINER_NAME")
	metricsPort := os.Getenv("METRICS_PORT")
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor))
	bot_to_server_proto.RegisterCongratulationServiceServer(grpcServer, serverImpl)
	grpc_prometheus.Register(grpcServer)

	group := run.Group{}
	group.Add(func() error {
		server, err := net.Listen("tcp", address+port)
		if err != nil {
			return fmt.Errorf("couldn't have started grpc server: %s", err.Error())
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
	if err = group.Run(); err != nil {
		utils.Logger.Fatal("Error while running a group", zap.String("error", err.Error()))
	}

}

func GRPCClientConfiguration() gen_proto.CongratulationServiceClient {
	port := os.Getenv("GRPC_PORT")
	name := os.Getenv("AI_CALLER_CONTAINER_NAME")
	conn, err := grpc.Dial(name+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.Logger.Fatal("Can't connect to another container in order to start communication", zap.String("error", err.Error()))
	}
	client := gen_proto.NewCongratulationServiceClient(conn)

	return client
}

func dbConfiguration() *gorm.DB {
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_port := os.Getenv("DB_PORT")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_CONTAINER_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			localization, err := time.LoadLocation("Europe/Moscow")
			if err != nil {
				utils.Logger.Fatal("Invalid time zone", zap.String("time zone", localization.String()))
			}
			return time.Now().In(localization)
		},
	})
	if err != nil {
		utils.Logger.Fatal("Can't open connection with the database", zap.String("error", err.Error()))
	}
	err = db.AutoMigrate(model.User{}, model.Chat{})
	if err != nil {
		utils.Logger.Fatal("Can't complete db migration", zap.String("error", err.Error()))
	}
	return db
}
