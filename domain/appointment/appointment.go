// Declares the package name
package appointment

//	Import library
import "time"

//	Declare Appointment Entity
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
