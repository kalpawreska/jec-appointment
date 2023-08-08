package handler

import (
	"errors"
	appointmentPB "jec-appointment/domain/appointment"
	"jec-appointment/model"
	"jec-appointment/pkg/dbconnect"
)

func NewDeleteAppointment(req *appointmentPB.AppointmentDeleteProto) (*appointmentPB.AppointmentDeleteResponseProto, error) {
	if req == nil {
		return nil, errors.New("empty request")
	}

	request := req.Deleteappointment
	if request == nil {
		return nil, errors.New("incomplete parameter")
	}
	if len(request.AppointmentNo) == 0 && len(request.HealthcareId) == 0 {
		return nil, errors.New("incomplete parameter")
	}

	db := dbconnect.DB
	appointment := model.Appointment{}

	prep := db.Unscoped().Model(&model.Appointment{})
	if len(request.AppointmentNo) > 0 {
		prep = prep.Where("appoint_no = ?", request.AppointmentNo)
	}
	if len(request.HealthcareId) > 0 {
		prep = prep.Where("healthcare_id = ?", request.HealthcareId)
	}
	err := prep.Delete(&appointment).Error

	resp := appointmentPB.AppointmentDeleteResponseProto{
		HealthcareId:  request.HealthcareId,
		AppointmentNo: request.AppointmentNo,
	}

	return &resp, err
}
