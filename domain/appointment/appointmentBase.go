// Declares the package name
package appointment

//	Import library
import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

func RouterInit(r fiber.Router) {
	svc := AppointmentService()
	handler := AppointmentHandler(svc)

	appointmentApi := r.Group("/api")
	appointmentApi.Post("/appointment", handler.Create)
}

func RouterInitWithDB(r fiber.Router, dbx *sqlx.DB) {
	svc := AppointmentService()
	handler := AppointmentHandler(svc)

	appointmentApi := r.Group("/api")
	appointmentApi.Post("/appointment", handler.Create)
}
