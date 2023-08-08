package main

import (
	"fmt"
	"log"
	"net"

	"jec-appointment/config"
	appointmentPB "jec-appointment/domain/appointment"
	"jec-appointment/pkg/dbconnect"
	"jec-appointment/pkg/utils"
	"jec-appointment/service"

	"google.golang.org/grpc"
)

func init() {
	utils.LoadEnvironment(config.Environment)
	dbconnect.InitDatabase()
}

// @title JEC Appointment Go gRPC Service
// @version 1.0.0
// @description JEC Appointment Go gRPC Service
// @contact.name JEC
// @contact.email no-reply@jec.co.id
// @host jec.co.id:50051
// @schemes https
// @BasePath /grpc
func main() {
	// Start the gRPC server
	grpcServer := grpc.NewServer()
	appointmentPB.RegisterAppointmentServiceServer(grpcServer, &service.AppointmentService{})

	port := utils.GetEnv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("gRPC server listening on :%s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
