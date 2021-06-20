package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"linkShortener/internal/pkg/shortLink/delivery"
	"linkShortener/internal/pkg/shortLink/repository"
	"linkShortener/internal/pkg/shortLink/usecase"
	proto_linkShortener "linkShortener/proto"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ConfigPath = "config.json"

type Config struct {
	Host string `json:"host"`
	Port int `json:"port"`

	DBHost string `json:"dbHost"`
	DBUser string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBName string `json:"dbName"`
	DBPort int `json:"dbPort"`

	Alphabet string `json:"alphabet"`
	LenShortenLink int `json:"len_shorten_link"`
}
func main() {
	var config Config
	configFile, err := os.Open(ConfigPath)
	if err != nil {
		log.Fatalf("cannot open config file: %s", err.Error())
	}

	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatalf("cannot unmarshal config file: %s", err.Error())
	}
	configFile.Close()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect to databse: %v", err.Error())
	}
	postgresDB, err := db.DB()
	if err != nil {
		log.Fatalf("cannot connect to databse: %v", err.Error())
	}
	defer postgresDB.Close()
	err = postgresDB.Ping()
	if err != nil {
		log.Fatalf("cannot connect to databse: %v", err.Error())
	}

	server := &delivery.ShortLinkServer{
		&usecase.ShortLinkUC{
			&repository.ShortLinkRep{
				db,
			},
			[]byte(config.Alphabet),
			config.LenShortenLink,
		},
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto_linkShortener.RegisterLinkShortenerServer(grpcServer, server)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("INFO: TCP server has started at %s\n", addr)
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-quit
	log.Println("Interrupt signal received. Shutting down server...")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	grpcServer.GracefulStop()
}

