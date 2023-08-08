package handler

import (
	"errors"
	appointmentPB "jec-appointment/domain/appointment"
	"jec-appointment/model"
	"jec-appointment/pkg/dbconnect"
	"jec-appointment/pkg/utils"

	"github.com/go-openapi/strfmt"
)

type UpdateAppointmentHandler struct {
	Request  *appointmentPB.AppointmentUpdateProto
	Response *appointmentPB.AppointmentProto
}

func NewUpdateAppointmentHandler(req *appointmentPB.AppointmentUpdateProto) (*UpdateAppointmentHandler, error) {
	pa := UpdateAppointmentHandler{
		Request: req,
	}

	if errorMsg := pa.ValidateRequest(); errorMsg != nil {
		return nil, errors.New(*errorMsg)
	}
	return &pa, nil
}

func (pa *UpdateAppointmentHandler) ValidateRequest() (errorMsg *string) {
	if pa.Request == nil {
		return utils.Strptr("Empty request")
	}
	req := pa.Request.Updateappointment
	if req == nil {
		return utils.Strptr("Empty request")
	}

	// Validate empty request
	if len(req.AppointmentNo) == 0 || req.AppointmentDate == nil || req.AppointmentTime == nil || len(req.HealthcareId) == 0 || len(req.ParamedicId) == 0 || len(req.PatientId) == 0 {
		return utils.Strptr("All required parameter need to be filled")
	}

	// Validate appointment date format
	if req.AppointmentDate.CheckValid() != nil {
		return utils.Strptr("Invalid appointment date format")
	}

	// Validate appointment time format
	if req.AppointmentTime.CheckValid() != nil {
		return utils.Strptr("Invalid appointment time format")
	}

	return nil
}

func (pa *UpdateAppointmentHandler) Update() error {
	db := dbconnect.DB

	request := pa.Request.Updateappointment

	appointmentDate := strfmt.Date(request.AppointmentDate.AsTime())
	appointmentTime := request.AppointmentTime.AsTime().Format("15:04:00")
	data := model.Appointment{
		AppointmentNo:   &request.AppointmentNo,
		HealthcareID:    &request.HealthcareId,
		ParamedicID:     &request.ParamedicId,
		PatientID:       &request.PatientId,
		AppointmentDate: &appointmentDate,
		AppointmentTime: &appointmentTime,
		ScheduleSlotID:  utils.Intptr(int(request.ScheduleSlotId)),
		UserCreate:      &request.UserCreate,
	}

	mapUpdate := map[string]interface{}{}
	if len(request.HealthcareId) > 0 {
		mapUpdate["healthcare_id"] = request.HealthcareId
	}
	if len(request.ParamedicId) > 0 {
		mapUpdate["paramedic_id"] = request.ParamedicId
	}
	if len(request.PatientId) > 0 {
		mapUpdate["patient_id"] = request.PatientId
	}
	if request.AppointmentDate != nil {
		mapUpdate["appointment_date"] = request.AppointmentDate.AsTime().Format("2006-01-02")
	}
	if request.AppointmentTime != nil {
		mapUpdate["appointment_time"] = request.AppointmentTime.AsTime().Format("15:04:05")
	}
	if request.ScheduleSlotId > 0 {
		mapUpdate["schedule_slot_id"] = request.ScheduleSlotId
	}
	if len(request.UserCreate) > 0 {
		mapUpdate["user_create"] = request.UserCreate
	}

	if len(mapUpdate) == 0 {
		return errors.New("no updated data")
	}

	err := db.Where("appoint_no = ?", request.AppointmentNo).Updates(&mapUpdate).Error
	if err != nil {
		return err
	}

	var response appointmentPB.AppointmentProto
	utils.Merge(data, &response)
	pa.Response = &response

	return nil
}
