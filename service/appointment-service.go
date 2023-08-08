package service

import (
	"context"
	appointmentPB "jec-appointment/domain/appointment"
	"jec-appointment/service/handler"

	"google.golang.org/protobuf/types/known/emptypb"
)

type AppointmentService struct {
	appointmentPB.UnimplementedAppointmentServiceServer
}

// Add index page
// @Summary Post appointment data to database
// @Description Post appointment data to database
// @Accept  application/json
// @Produce  application/json
// @Param data body appointmentPB.AppointmentAddProto true "Albedo payload"
// @Success 200 {object} appointmentPB.AppointmentProto "success"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal error"
// @Router /Appointment/Add [post]
// @Tags Appointment
func (a *AppointmentService) Add(ctx context.Context, req *appointmentPB.AppointmentAddProto) (*appointmentPB.AppointmentProto, error) {
	handler, err := handler.NewAddAppointmentHandler(req)
	if err != nil {
		return nil, err
	}
	err = handler.Store()
	if err != nil {
		return nil, err
	}
	return handler.Response, nil
}

// Get index page
// @Summary Get single appointment based on appoint_no
// @Description Get single appointment based on appoint_no
// @Accept  application/json
// @Produce  application/json
// @Param data body appointmentPB.AppointmentGetProto true "Albedo payload"
// @Success 200 {object} appointmentPB.AppointmentProto "success"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal error"
// @Router /Appointment/Get [post]
// @Tags Appointment
func (a *AppointmentService) Get(ctx context.Context, req *appointmentPB.AppointmentGetProto) (*appointmentPB.AppointmentProto, error) {
	resp, err := handler.NewGetAppointment(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// List index page
// @Summary Get lists appointment based on filter
// @Description Get lists appointment based on filter
// @Accept  application/json
// @Produce  application/json
// @Param data body appointmentPB.AppointmentProto true "Albedo payload"
// @Success 200 {object} appointmentPB.AppointmentListProto "success"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal error"
// @Router /Appointment/List [post]
// @Tags Appointment
func (a *AppointmentService) List(ctx context.Context, req *emptypb.Empty) (*appointmentPB.AppointmentListProto, error) {
	handler, err := handler.NewListAppointment(nil)
	if err != nil {
		return nil, err
	}
	data := handler.GetData()
	return data, nil
}

// ListByParam index page
// @Summary Get lists appointment based on filter
// @Description Get lists appointment based on filter
// @Accept  application/json
// @Produce  application/json
// @Param data body appointmentPB.AppointmentProto true "Albedo payload"
// @Success 200 {object} appointmentPB.AppointmentListProto "success"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal error"
// @Router /Appointment/ListByParam [post]
// @Tags Appointment
func (a *AppointmentService) ListByParam(ctx context.Context, req *appointmentPB.AppointmentProto) (*appointmentPB.AppointmentListProto, error) {
	handler, err := handler.NewListAppointment(req)
	if err != nil {
		return nil, err
	}
	data := handler.GetData()
	return data, nil
}

// Update index page
// @Summary Update appointment based on filter
// @Description Update appointment based on filter
// @Accept  application/json
// @Produce  application/json
// @Param data body appointmentPB.AppointmentUpdateProto true "Albedo payload"
// @Success 200 {object} []appointmentPB.AppointmentProto "success"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal error"
// @Router /Appointment/Update [post]
// @Tags Appointment
func (a *AppointmentService) Update(ctx context.Context, req *appointmentPB.AppointmentUpdateProto) (*appointmentPB.AppointmentProto, error) {
	handler, err := handler.NewUpdateAppointmentHandler(req)
	if err != nil {
		return nil, err
	}
	err = handler.Update()
	if err != nil {
		return nil, err
	}
	return handler.Response, nil
}

// Delete index page
// @Summary Delete appointment based on filter
// @Description Delete appointment based on filter
// @Accept  application/json
// @Produce  application/json
// @Param data body appointmentPB.AppointmentDeleteProto true "Albedo payload"
// @Success 200 {object} appointmentPB.AppointmentDeleteResponseProto "success"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal error"
// @Router /Appointment/Delete [post]
// @Tags Appointment
func (a *AppointmentService) Delete(ctx context.Context, req *appointmentPB.AppointmentDeleteProto) (*appointmentPB.AppointmentDeleteResponseProto, error) {
	data, err := handler.NewDeleteAppointment(req)
	return data, err
}
