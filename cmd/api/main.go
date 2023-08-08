// Declares the package name
package main

//	Import library
import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kalpawreska/jec-appointment/domain/appointment"
	customvalidator "github.com/kalpawreska/jec-appointment/pkg/jecvalidator"
)

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

func main() {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: customvalidator.HttpErrorHandler,
		},
	)

	appointment.RouterInit(app)

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
