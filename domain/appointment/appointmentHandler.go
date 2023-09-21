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
	// Access Service
	appointmentSvc appointmentService
}

// Declare Appointment Handler
// @Param appointmentService aka p_oService
func NewAppointmentHandler(p_oService appointmentService) appointmentHandler {
	return appointmentHandler{
		appointmentSvc: p_oService,
	}
}

// Get Appointment List.
// @Summary List of Appointment Object.
// @Description Get Appointment List.
// @Accept application/json
// @Produce application/json
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 404 {object} map[string]interface{} "no found"
// @Router /appointment [GET]
// @Tags Appointment
func (h appointmentHandler) List(ctx *fiber.Ctx) error {
	resp, err := h.appointmentSvc.ListService(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"type":    "https://tools.ietf.org/html/rfc7231#section-6.5.1",
			"traceid": "-99",
			"title":   "Bad Request",
			"status":  fiber.StatusBadRequest,
			"detail":  err,
		})
	}

	if len(resp) == 0 {
		return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
			"traceid": "traceid",
			"title":   "No Data",
			"status":  fiber.StatusOK,
			"detail":  "Appointment data doesn't exists.",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"status":  fiber.StatusOK,
		"message": "Get Appointment List has been successfully!",
		"data":    &resp,
	})
}

// Get Appointment By Healthcare ID and Appointment No.
// @Summary Appointment Object.
// @Description Get Appointment By Healthcare ID and Appointment No.
// @Accept application/json
// @Produce application/json
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 404 {object} map[string]interface{} "no found"
// @Param healthcare_id query string false "Healthcare ID"
// @Param appointment_no query string false "Appointment No"
// @Router /appointment/get [GET]
// @Tags Appointment
func (h appointmentHandler) Get(ctx *fiber.Ctx) error {
	req := new(AppointmentRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}
	if err := ValidateStruct(req); err != nil {
		return ctx.JSON(err)
	}

	resp, err := h.appointmentSvc.GetService(ctx.Context(), *req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"type":    "https://tools.ietf.org/html/rfc7231#section-6.5.1",
			"traceid": "-99",
			"title":   "Bad Request",
			"status":  fiber.StatusBadRequest,
			"detail":  err,
		})
	}

	if len(resp) == 0 {
		return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
			"traceid": "traceid",
			"title":   "No Data",
			"status":  fiber.StatusOK,
			"detail":  "Appointment data doesn't exists.",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"status":  fiber.StatusOK,
		"message": "Get Apppointment has been successfully!",
		"data":    &resp,
	})
}

// Create Appointment.
// @Summary A newly created Appointment.
// @Description Create Appointment.
// @Accept application/json
// @Produce application/json
// @Param data body AppointmentRequest true "Data request"
// @Success 201 {object} map[string]interface{} "created"
// @Failure 204 {object} map[string]interface{} "no content"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Router /appointment [POST]
// @Tags Appointment
func (h appointmentHandler) Create(ctx *fiber.Ctx) error {
	req := new(AppointmentRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}
	if err := ValidateStruct(req); err != nil {
		return ctx.JSON(err)
	}

	err := h.appointmentSvc.CreateService(ctx.Context(), *req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"type":    "https://tools.ietf.org/html/rfc7231#section-6.5.1",
			"traceid": "-99",
			"title":   "Bad Request",
			"status":  fiber.StatusBadRequest,
			"detail":  err,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"status":  fiber.StatusCreated,
		"message": "Create appointment has been successfully!",
		"data":    &req,
	})
}

// Update Appointment.
// @Summary A newly updated Appointment.
// @Description Update Appointment.
// @Accept application/json
// @Produce application/json
// @Param data body AppointmentUpdateProto true "Data request"
// @Success 200 {object} map[string]interface{} "success"
// @Success 201 {object} map[string]interface{} "created"
// @Failure 204 {object} map[string]interface{} "no content"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 404 {object} map[string]interface{} "no found"
// @Router /appointment [PUT]
// @Tags Appointment
func (h appointmentHandler) Update(ctx *fiber.Ctx) error {
	req := new(AppointmentUpdateProto)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}
	if err := ValidateStruct(req); err != nil {
		return ctx.JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"status":  fiber.StatusOK,
		"message": "Updated appointment has been successfully!",
		"data":    &req,
	})
}

// Delete Appointment.
// @Summary A newly deleted Appointment.
// @Description Delete Appointment.
// @Accept application/json
// @Produce application/json
// @Param data body AppointmentDeleteProto true "Data request"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 202 {object} map[string]interface{} "accepted"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 404 {object} map[string]interface{} "no found"
// @Router /appointment [DELETE]
// @Tags Appointment
func (h appointmentHandler) Delete(ctx *fiber.Ctx) error {
	req := new(AppointmentDeleteProto)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}
	if err := ValidateStruct(req); err != nil {
		return ctx.JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"status":  fiber.StatusOK,
		"message": "Delete appointment has been successfully!",
		"data":    &req,
	})
}
