// Declares the package name
package appointment

//	Import library
import (
	"github.com/gofiber/fiber/v2"
)

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

// Declare Appointment Handler construct
type appointmentHandler struct {
	appointmentSvc appointmentService
}

// Declare Create Appointment Handler
func AppointmentHandler(p_oService appointmentService) appointmentHandler {
	return appointmentHandler{
		appointmentSvc: p_oService,
	}
}

func (h appointmentHandler) Create(ctx *fiber.Ctx) error {
	req := new(AppointmentRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}
	if err := ValidateStruct(req); err != nil {
		return ctx.JSON(err)
	}

	err := h.appointmentSvc.CreateService(*req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"status":  fiber.StatusCreated,
		"message": "Create appointment has been successfully!",
		"data":    &req,
	})
}
