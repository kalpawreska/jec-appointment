package handler

import (
	"errors"
	appointmentPB "jec-appointment/domain/appointment"
	"jec-appointment/model"
	"jec-appointment/pkg/dbconnect"
	"jec-appointment/pkg/utils"
)

func NewGetAppointment(req *appointmentPB.AppointmentGetProto) (*appointmentPB.AppointmentProto, error) {
	var response appointmentPB.AppointmentProto
	if req == nil {
		return nil, errors.New("empty request")
	}

	if len(req.AppointmentNo) == 0 && len(req.HealthcareId) == 0 {
		return nil, errors.New("incomplete parameter")
	}

	db := dbconnect.DB
	appointment := model.Appointment{}
	db.Model(&model.Appointment{}).Where("appoint_no = ?", req.AppointmentNo).Take(&appointment)

	if appointment.AppointmentNo == nil {
		return nil, errors.New("appointment not found")
	}

	utils.Merge(appointment, &response)
	return &response, nil
}
