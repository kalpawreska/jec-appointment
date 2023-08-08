package handler

import (
	appointmentPB "jec-appointment/domain/appointment"
	"jec-appointment/model"
	"jec-appointment/pkg/dbconnect"
	"jec-appointment/pkg/utils"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetListAppointmentHandler struct {
	Request  *appointmentPB.AppointmentProto
	Response *appointmentPB.AppointmentListProto
}

func NewListAppointment(req *appointmentPB.AppointmentProto) (*GetListAppointmentHandler, error) {
	ga := GetListAppointmentHandler{
		Request: req,
	}

	return &ga, nil
}

func (ga *GetListAppointmentHandler) GetData() *appointmentPB.AppointmentListProto {
	var collection appointmentPB.AppointmentListProto

	request := ga.Request

	db := dbconnect.DB
	appointments := []model.Appointment{}

	prep := db.Model(&model.Appointment{})
	if request != nil {
		if len(request.AppointmentNo) > 0 {
			prep = prep.Where("appoint_no = ?", request.AppointmentNo)
		}
		if len(request.HealthcareId) > 0 {
			prep = prep.Where("healthcare_id = ?", request.HealthcareId)
		}
		if len(request.ParamedicId) > 0 {
			prep = prep.Where("paramedic_id = ?", request.ParamedicId)
		}
		if len(request.PatientId) > 0 {
			prep = prep.Where("patient_id = ?", request.PatientId)
		}
	}
	prep.Order("create_at DESC").Find(&appointments)

	apptMetas := []*appointmentPB.AppointmentProto{}
	for _, appt := range appointments {
		apptMeta := appointmentPB.AppointmentProto{}
		utils.Merge(appt, &apptMeta)
		apptMeta.CreateAt = timestamppb.New(time.Time(*appt.CreateAt))

		apptMetas = append(apptMetas, &apptMeta)
	}
	collection.Appointments = apptMetas

	return &collection
}
