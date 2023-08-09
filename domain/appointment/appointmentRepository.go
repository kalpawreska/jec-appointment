// Declares the package name
package appointment

//	Import library
import (
	"context"
	tools "jec-appointment/pkg/jectools"

	"github.com/jmoiron/sqlx"
)

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

func (r appointmentRepository) GetListAppointment(ctx context.Context, limit, page int) ([]Appointment, int, error) {
	var (
		result = make([]Appointment, 0)
		offset = (page * limit) - limit
	)

	query := `SELECT * FROM appointments LIMIT ? OFFSET ?`
	query = r.db.Rebind(query)

	// fmt.Println(query)
	// fmt.Println(limit, offset)
	// fmt.Println("...............")

	if err := r.db.SelectContext(ctx, &result, query, limit, offset); err != nil {
		return result, 0, err
	}

	var total int
	queryCount := `SELECT COUNT(*) AS total FROM appointments`
	if err := r.db.GetContext(ctx, &total, queryCount); err != nil {
		return result, 0, err
	}

	return result, total, nil
}

func (r appointmentRepository) GetByAppointmentNo(ctx context.Context, appointmentNo, healthCareID string) ([]Appointment, error) {
	var (
		res  []Appointment
		args []any
	)

	query := `SELECT * FROM appointments WHERE healthcare_id = ?`
	args = append(args, healthCareID)

	if appointmentNo != "" {
		query += ` AND appointment_no = ?`
		args = append(args, appointmentNo)
	}

	query = r.db.Rebind(query)

	if err := r.db.SelectContext(ctx, &res, query, args...); err != nil {
		return res, err
	}

	return res, nil
}
