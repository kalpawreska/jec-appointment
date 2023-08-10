// Declares the package name
package appointment

//	Import library
import "time"

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

//	Declare Appointment Entity
type Appointment struct {
	HealthcareId    string
	AppointmentNo   string
	ParamedicId     string
	PatientId       string
	AppointmentDate time.Time
	AppointmentTime time.Time
	IsVoid          bool
	UserCreate      string
	CreateAt        time.Time
	ScheduleSlotId  int
}
