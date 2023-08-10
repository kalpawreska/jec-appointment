// Declares the package name
package appointment

//	Import library
import (
	"errors"
	"time"
)

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

// [Response] Declare Appointment List construct
type AppointmentListResponse struct {
	HealthcareId    string    `json:"HealthcareId"`
	AppointmentNo   string    `json:"AppointmentNo"`
	ParamedicId     string    `json:"ParamedicId"`
	PatientId       string    `json:"PatientId"`
	AppointmentDate time.Time `json:"AppointmentDate"`
	AppointmentTime time.Time `json:"AppointmentTime"`
	IsVoid          bool      `json:"IsVoid"`
	UserCreate      string    `json:"UserCreate"`
	CreateAt        time.Time `json:"CreateAt"`
	ScheduleSlotId  int       `json:"ScheduleSlotId"`
}

// [Request] Declare Appointment List construct
type AppointmentRequest struct {
	HealthcareId    string    `json:"HealthcareId"`
	AppointmentNo   string    `json:"AppointmentNo"`
	ParamedicId     string    `json:"ParamedicId"`
	PatientId       string    `json:"PatientId"`
	AppointmentDate time.Time `json:"AppointmentDate"`
	AppointmentTime time.Time `json:"AppointmentTime"`
	IsVoid          bool      `json:"IsVoid"`
	UserCreate      string    `json:"UserCreate"`
	CreateAt        time.Time `json:"CreateAt"`
	ScheduleSlotId  int       `json:"ScheduleSlotId"`
}

//  #region Method

// Convert Object to Entity
func (a AppointmentRequest) ParseToEntity() Appointment {
	return Appointment{
		HealthcareId:    a.HealthcareId,
		AppointmentNo:   a.AppointmentNo,
		ParamedicId:     a.ParamedicId,
		PatientId:       a.PatientId,
		AppointmentDate: a.AppointmentDate,
		AppointmentTime: a.AppointmentTime,
		IsVoid:          false,
		UserCreate:      a.UserCreate,
		CreateAt:        time.Now(),
		ScheduleSlotId:  a.ScheduleSlotId,
	}
}

// Validate Data Request
func (a AppointmentRequest) Validate() (err error) {
	if len(a.HealthcareId) < 3 {
		return errors.New("appointment Healthcare ID minimum 3 length")
	}
	if len(a.AppointmentNo) < 14 {
		return errors.New("appointment Appointment No minimum 14 length")
	}
	if len(a.ParamedicId) < 5 {
		return errors.New("appointment Paramedic ID minimum 5 length")
	}
	if len(a.PatientId) < 13 {
		return errors.New("appointment Patient ID minimum 13 length")
	}
	if len(a.UserCreate) < 5 {
		return errors.New("appointment User Create No minimum 5 length")
	}

	return nil
}

// #endregion Method
