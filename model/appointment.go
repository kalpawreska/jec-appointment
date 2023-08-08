package model

import "github.com/go-openapi/strfmt"

type Appointment struct {
	AppointmentNo   *string          `json:"appointment_no" gorm:"column:appoint_no;type:varchar(20);primaryKey"`
	HealthcareID    *string          `json:"healthcare_id" gorm:"type:varchar(3);not null"`
	ParamedicID     *string          `json:"paramedic_id" gorm:"type:varchar(10);not null"`
	PatientID       *string          `json:"patient_id" gorm:"type:varchar(30);not null"`
	AppointmentDate *strfmt.Date     `json:"appointment_date" gorm:"type:date;not null"`
	AppointmentTime *string          `json:"appointment_time" gorm:"type:time;not null"`
	IsVoid          *bool            `json:"is_void" gorm:"type:bool"`
	UserCreate      *string          `json:"user_create" gorm:"type:varchar(10)"`
	CreateAt        *strfmt.DateTime `json:"create_at" gorm:"type:timestamp"`
	ScheduleSlotID  *int             `json:"schedule_slot_id"`
}

func (Appointment) TableName() string {
	return "appointments"
}
