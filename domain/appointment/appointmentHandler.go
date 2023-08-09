// Declares the package name
package appointment

//	Import library
import (
	"math"

	"github.com/gofiber/fiber/v2"
)

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
// @Success 200 {object} "success"
// @Failure 400 {object} "bad request"
// @Failure 404 {object} "no found"
// @Router /appointment [GET]
// @Tags Appointment
func (h appointmentHandler) List(ctx *fiber.Ctx) error {
	var (
		limit int = ctx.QueryInt("limit")
		page  int = ctx.QueryInt("page")
	)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 5
	}

	items, total, err := h.appointmentSvc.ListAppointment(ctx.UserContext(), limit, page)
	if err != nil {
		return err
	}

	totalPage := math.Ceil(float64(total) / float64(limit))
	if totalPage == 0 {
		totalPage = 1
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"status":     fiber.StatusOK,
		"message":    "Get Appointment List has been successfully!",
		"data":       items,
		"total":      total,
		"total_page": totalPage,
	})
}

// Get Appointment By Healthcare ID and Appointment No.
// @Summary Appointment Object.
// @Description Get Appointment By Healthcare ID and Appointment No.
// @Accept application/json
// @Produce application/json
// @Success 200 {object} "success"
// @Failure 400 {object} "bad request"
// @Failure 404 {object} "no found"
// @Router /appointment/get?{healthcare_id}&{appointment_no} [GET]
// @Tags Appointment
func (h appointmentHandler) Get(ctx *fiber.Ctx) error {
	var (
		appointmentNo string = ctx.Query("appointment_no")
		healthCareID  string = ctx.Query("healthcare_id")
	)

	if healthCareID == "" {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(map[string]string{
			"message": "please input healthcare id",
		})
	}

	res, err := h.appointmentSvc.DetailAppointment(ctx.UserContext(), appointmentNo, healthCareID)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"status":  fiber.StatusOK,
		"message": "Get Appointment has been successfully!",
		"data":    res,
	})
}

// Create Appointment.
// @Summary A newly created Appointment.
// @Description Create Appointment.
// @Accept application/json
// @Produce application/json
// @Success 201 {object} "created"
// @Failure 204 {object} "no content"
// @Failure 400 {object} "bad request"
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
		return err
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
// @Success 200 {object} "success"
// @Success 201 {object} "created"
// @Failure 204 {object} "no content"
// @Failure 400 {object} "bad request"
// @Failure 404 {object} "no found"
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
// @Success 200 {object} "success"
// @Failure 202 {object} "accepted"
// @Failure 400 {object} "bad request"
// @Failure 404 {object} "no found"
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
