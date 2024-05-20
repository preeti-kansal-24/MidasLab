package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	handlergrpc "preeti-kansal-24/MidasLab.git/application/grpc"
	"preeti-kansal-24/MidasLab.git/application/repository"
	"preeti-kansal-24/MidasLab.git/application/service"
	"preeti-kansal-24/MidasLab.git/application/transformer"
	"preeti-kansal-24/MidasLab.git/constants"
	"preeti-kansal-24/MidasLab.git/database"
	"preeti-kansal-24/MidasLab.git/kafka"
	proto "preeti-kansal-24/MidasLab.git/proto/src/go"
	"syscall"
)

var (
	port         = flag.Int("port", 50051, "The server port")
	otpTopic     = flag.String("otp-topic", "generate-otp-topic", "It will asynchronously generate otp while user is getting registered")
	twilioSid    = flag.String("twilio-sid", "", "provide this through program line args while running main.go")
	twilioAuth   = flag.String("twilio-auth", "", "provide this through program line args while running main.go")
	twilioFromPh = flag.String("twilio-from-phone-number", "", "provide a valid twilio from-phone-number")
)

func main() {
	flag.Parse()
	setupConfigs()

	//create db conn and migrate all schemas
	dbConn := database.CreateDBConn("postgresql://user:midas@localhost:25432/midas?sslmode=disable", true)
	database.Migrate()

	//Registering repositories
	authStore := repository.NewAuthStore(dbConn)
	otpStore := repository.NewOtpStore(dbConn)

	//Registering services
	authService := service.NewAuthService(authStore, otpStore)
	otpService := service.NewOtpService(authStore, otpStore)

	//Registering transformers
	authTransformer := transformer.NewAuthTransformer()

	//Registering handlers
	authHandler := handlergrpc.NewAuthHandler(authService, authTransformer)
	otpHandler := handlergrpc.NewOtpHandler(otpService)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	producer := kafka.CreateNewProducer()

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Failed to close producer: %v", err)
		}
	}()

	fmt.Println("otp topic:", constants.GenerateOtpTopic)
	_, err := kafka.NewConsumer(ctx, constants.GenerateOtpTopic, otpStore)
	if err != nil {
		fmt.Println("Failed to create consumer:", err)
		return
	}

	// Handle signals for graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	// Run gRPC server
	go runGRPCServer(authHandler, otpHandler, port)
	// Wait for shutdown signal
	<-sigchan
	fmt.Println("Shutting down gracefully...")
	cancel()
}

func setupConfigs() {
	constants.GenerateOtpTopic = *otpTopic
	constants.TwilioSID = *twilioSid
	constants.TwilioAuth = *twilioAuth
	constants.TwilioFromPhone = *twilioFromPh
}

func runGRPCServer(authHandler proto.AuthServiceServer, otpHandler proto.OtpServiceServer, port *int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterAuthServiceServer(s, authHandler)
	proto.RegisterOtpServiceServer(s, otpHandler)
	log.Printf("server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
