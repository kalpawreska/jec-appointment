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
	// svc := AppointmentService()
	handler := AppointmentHandler()

	appointmentApi := r.Group("/api")
	appointmentApi.Get("/appointment", handler.List)
	appointmentApi.Get("/appointment/get", handler.Get)
	appointmentApi.Post("/appointment", handler.Create)
	appointmentApi.Put("/appointment", handler.Update)
	appointmentApi.Delete("/appointment", handler.Delete)
}

func RouterInitWithDB(r fiber.Router, dbx *sqlx.DB) {
	// svc := AppointmentService()
	handler := AppointmentHandler()

	appointmentApi := r.Group("/api")
	appointmentApi.Post("/appointment", handler.Create)
}
