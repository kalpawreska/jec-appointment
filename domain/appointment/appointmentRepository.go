// Declares the package name
package appointment

//	Import library
import (
	"context"
	"errors"
	"log"

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

// Implement List Appointment Repository on Service
func (r appointmentRepository) ListRepo(ctx context.Context) ([]Appointment, error) {
	var (
		vrResult = make([]Appointment, 0)
	)

	query := `SELECT * FROM appointments
        WHERE is_void = false
        ORDER BY create_at DESC`
	query = r.db.Rebind(query)

	if err := r.db.SelectContext(ctx, &vrResult, query); err != nil {
		return vrResult, err
	}

	return vrResult, nil
}

// Implement Get Appointment Repository on Service
func (r appointmentRepository) GetRepo(ctx context.Context, p_strHealthcareId string, p_strAppointmentNo string) ([]Appointment, error) {
	var (
		vrResult []Appointment
		args     []any
	)

	if len(p_strHealthcareId) == 0 && len(p_strAppointmentNo) == 0 {
		return nil, errors.New("incomplete Parameter")
	}

	query := `SELECT * FROM appointments WHERE`
	// args = append(args, p_strHealthcareId)

	if p_strHealthcareId != "" {
		query += ` healthcare_id = ?`
		args = append(args, p_strHealthcareId)
	}
	if p_strAppointmentNo != "" {
		query += ` AND appointment_no = ?`
		args = append(args, p_strAppointmentNo)
	}
	// log.Println(query)
	query = r.db.Rebind(query)

	if err := r.db.SelectContext(ctx, &vrResult, query, args...); err != nil {
		return vrResult, err
	}

	log.Println(vrResult)
	return vrResult, nil
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
