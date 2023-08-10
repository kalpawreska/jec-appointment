package main

import (
	"fmt"
	"log"
	"net"

	"jec-appointment/domain/appointment"
	dbs "jec-appointment/pkg/jecconfiguration"
	tools "jec-appointment/pkg/jectools"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	//  Load Environtment
	err := godotenv.Load("././.env")
	if err != nil {
		log.Fatalf("Get Environtment Failed :%v", err)
	}

	//  Database Initial
	dbConn, err := dbs.ConnectSqlx(dbs.DbConfiguration{
		Host:       tools.GetEnv("POSTGRES_HOST"),
		Port:       tools.GetEnv("POSTGRES_PORT"),
		Dbname:     tools.GetEnv("POSTGRES_DBNAME"),
		Dbuser:     tools.GetEnv("POSTGRES_USER"),
		Dbpassword: tools.GetEnv("POSTGRES_PASSWORD"),
		Sslmode:    tools.GetEnv("POSTGRES_SSLMODE"),
	})
	if err != nil {
		panic(err)
	}

	if dbConn == nil {
		panic("Database [" + tools.GetEnv("POSTGRES_DBNAME") + "] Postgree Not Connected!")
	}
	log.Println("Database [" + tools.GetEnv("POSTGRES_DBNAME") + "] Postgree Connected!")

	srv := grpc.NewServer()
	appointment.RouterInitGRPC(srv, dbConn)

	log.Println("Registered GRPC Route ...")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", tools.GetEnv("GRPC_PORT")))
	if err != nil {
		panic(err)
	}

	log.Println("Appointment GRPC Server Running at port ", tools.GetEnv("GRPC_PORT"))
	err = srv.Serve(listen)
	if err != nil {
		panic(err)
	}
}
