// Declares the package name
package main

//	Import library
import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kalpawreska/jec-appointment/domain/appointment"
	dbs "github.com/kalpawreska/jec-appointment/pkg/jecconfiguration"
	tools "github.com/kalpawreska/jec-appointment/pkg/jectools"
	customvalidator "github.com/kalpawreska/jec-appointment/pkg/jecvalidator"
)

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

func main() {
	//  Load Environtment
	err := godotenv.Load("../../.env")
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

	//  Fiber Framework Initial
	app := fiber.New(
		fiber.Config{
			ErrorHandler: customvalidator.HttpErrorHandler,
		},
	)

	appointment.RouterInitWithDB(app, dbConn)

	log.Println("Appointment Services Running at port 9091")
	app.Listen(":9091")
}

/*
    ? OUTPUT :
    ? =====================================

{
    "data": {
        "HealthcareId": "002",
        "AppointmentNo": "AP230807-00001",
        "ParamedicId": "J0001",
        "PatientId": "KP202306-0001",
        "AppointmentDate": "2023-08-08T23:35:13+07:00",
        "AppointmentTime": "2023-08-08T23:35:13+07:00",
        "IsVoid": false,
        "UserCreate": "21239",
        "CreateAt": "2023-08-08T23:35:13+07:00",
        "ScheduleSlotId": 1
    },
    "message": "Create appointment has been successfully!",
    "status": 201
} */
