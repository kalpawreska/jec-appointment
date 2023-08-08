package handler

import (
	"errors"
	appointmentPB "jec-appointment/domain/appointment"
	"jec-appointment/model"
	"jec-appointment/pkg/dbconnect"
	"jec-appointment/pkg/utils"
	"time"

	"github.com/go-openapi/strfmt"
)

type PostAppointmentHandler struct {
	Request  *appointmentPB.AppointmentAddProto
	Response *appointmentPB.AppointmentProto
}

func NewAddAppointmentHandler(req *appointmentPB.AppointmentAddProto) (*PostAppointmentHandler, error) {
	pa := PostAppointmentHandler{
		Request: req,
	}

	if errorMsg := pa.ValidateRequest(); errorMsg != nil {
		return nil, errors.New(*errorMsg)
	}
	return &pa, nil
}

func (pa *PostAppointmentHandler) ValidateRequest() (errorMsg *string) {
	if pa.Request == nil {
		return utils.Strptr("Empty request")
	}
	req := pa.Request.Addappointment
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

	// Validate appoint_no duplicate or not
	if errMsg := validateDuplicateAppointment(req.AppointmentNo); errMsg != nil {
		return errMsg
	}

	// Validate same request or not
	if errMsg := validateSameRequestAppointment(req); errMsg != nil {
		return errMsg
	}

	return nil
}

func validateDuplicateAppointment(AppointmentNo string) *string {
	db := dbconnect.DB
	var duplicateCount int64
	db.Model(&model.Appointment{}).Where("appoint_no = ?", AppointmentNo).Count(&duplicateCount)

	if duplicateCount > 0 {
		return utils.Strptr("Appoint No " + AppointmentNo + " is already exists")
	}
	return nil
}

func validateSameRequestAppointment(req *appointmentPB.AppointmentProto) *string {
	db := dbconnect.DB
	var duplicateCount int64
	db.Model(&model.Appointment{}).
		Where("healthcare_id = ?", req.HealthcareId).
		Where("paramedic_id = ?", req.ParamedicId).
		Where("patient_id = ?", req.PatientId).
		Where("appointment_date = ?", req.AppointmentDate.AsTime().Format("2006-01-02")).
		Where("is_void = ?", false).
		Count(&duplicateCount)

	if duplicateCount > 0 {
		return utils.Strptr("Appointment with same request is already exists")
	}
	return nil
}

func (pa *PostAppointmentHandler) Store() error {
	db := dbconnect.DB

	request := pa.Request.Addappointment

	appointmentDate := strfmt.Date(request.AppointmentDate.AsTime())
	appointmentTime := request.AppointmentTime.AsTime().Format("15:04:05")
	data := model.Appointment{
		AppointmentNo:   &request.AppointmentNo,
		HealthcareID:    &request.HealthcareId,
		ParamedicID:     &request.ParamedicId,
		PatientID:       &request.PatientId,
		AppointmentDate: &appointmentDate,
		AppointmentTime: &appointmentTime,
		ScheduleSlotID:  utils.Intptr(int(request.ScheduleSlotId)),
		IsVoid:          utils.Boolptr(false),
		CreateAt:        utils.Datetimeptr(strfmt.DateTime(time.Now())),
		UserCreate:      &request.UserCreate,
	}
	err := db.Create(&data).Error
	if err != nil {
		return err
	}

	var response appointmentPB.AppointmentProto
	utils.Merge(data, &response)
	pa.Response = &response

	return nil
}
