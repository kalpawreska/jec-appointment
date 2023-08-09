// Declares the package name
package appointment

//	Import library
import (
	"context"

	"github.com/jmoiron/sqlx"
	tools "github.com/kalpawreska/jec-appointment/pkg/jectools"
)

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

// Declare Appointment Repository construct
type appointmentRepository struct {
	db *sqlx.DB
}

// Declare Appointment Repository
func NewAppointmentRepository(p_oDatabase *sqlx.DB) appointmentRepository {
	return appointmentRepository{db: p_oDatabase}
}

// Implements Add Appointment Repository on Service
func (r appointmentRepository) AddRepo(ctx context.Context, app Appointment) (p_strAppointmentNo string, err error) {
	query := `INSERT INTO appointments (
        healthcare_id, appointment_no, paramedic_id, patient_id, appointment_date, appointment_time, is_void, user_create, create_at, schedule_slot_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning appointment_no`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "-99", err
	}

	err = stmt.QueryRow(
		tools.NewSQLNullString(app.HealthcareId),
		tools.NewSQLNullString(app.AppointmentNo),
		tools.NewSQLNullString(app.ParamedicId),
		tools.NewSQLNullString(app.PatientId),
		app.AppointmentDate,
		app.AppointmentTime,
		app.IsVoid,
		tools.NewSQLNullString(app.UserCreate),
		app.CreateAt,
		app.ScheduleSlotId,
	).Scan(&p_strAppointmentNo)

	if err != nil {
		return "-99", err
	}

	return
}
