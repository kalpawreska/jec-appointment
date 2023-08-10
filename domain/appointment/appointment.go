// Declares the package name
package appointment

//	Import library
import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

// Declare Appointment Entity
type Appointment struct {
	HealthcareId    string    `db:"healthcare_id"`
	AppointmentNo   string    `db:"appointment_no"`
	ParamedicId     string    `db:"paramedic_id"`
	PatientId       string    `db:"patient_id"`
	AppointmentDate time.Time `db:"appointment_date"`
	AppointmentTime time.Time `db:"appointment_time"`
	IsVoid          bool      `db:"is_void"`
	UserCreate      string    `db:"user_create"`
	CreateAt        time.Time `db:"create_at"`
	ScheduleSlotId  int       `db:"schedule_slot_id"`
}

func (a *Appointment) ParseToAppointmentProto() *AppointmentProto {
	ap := AppointmentProto{
		AppointmentNo:   a.AppointmentNo,
		HealthcareId:    a.HealthcareId,
		ParamedicId:     a.ParamedicId,
		PatientId:       a.PatientId,
		AppointmentDate: timestamppb.New(a.AppointmentDate),
		AppointmentTime: timestamppb.New(a.AppointmentTime),
		IsVoid:          false,
		UserCreate:      a.UserCreate,
		CreateAt:        timestamppb.New(a.CreateAt),
		ScheduleSlotId:  int64(a.ScheduleSlotId),
	}

	return &ap
}
