// Declares the package name
package appointment

//	Import library
import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	tools "jec-appointment/pkg/jectools"

	"github.com/jmoiron/sqlx"
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

func (r appointmentRepository) GetSingleRepo(ctx context.Context, appointmentNo string, healthcareID string) (*Appointment, error) {
	if len(appointmentNo) == 0 && len(healthcareID) == 0 {
		return nil, errors.New("incomplete parameter")
	}

	db := r.db
	appointment := Appointment{}
	query := "SELECT * FROM appointments WHERE"
	param := []interface{}{}

	if len(appointmentNo) > 0 {
		query = query + " appointment_no = $1"
		param = append(param, appointmentNo)
	}
	if len(healthcareID) > 0 {
		a := "$1"
		if len(appointmentNo) > 0 {
			a = "$2"
			query = query + " AND "
		}
		query = query + " healthcare_id = " + a
		param = append(param, healthcareID)
	}

	db.Get(&appointment, query, param...)
	if len(appointment.AppointmentNo) == 0 {
		return nil, errors.New("appointment not found")
	}

	return &appointment, nil
}

func (r appointmentRepository) ListRepo(ctx context.Context, request *AppointmentRequest) (*[]Appointment, error) {
	db := r.db
	appointments := []Appointment{}

	qry := "SELECT * FROM appointments "
	where := []string{}
	params := []interface{}{}
	iterate := 1
	if request != nil {
		if len(request.AppointmentNo) > 0 {
			where = append(where, fmt.Sprintf("appointment_no = $%d", iterate))
			params = append(params, request.AppointmentNo)
			iterate = iterate + 1
		}
		if len(request.HealthcareId) > 0 {
			where = append(where, fmt.Sprintf("healthcare_id = $%d", iterate))
			params = append(params, request.HealthcareId)
			iterate = iterate + 1
		}
		if len(request.ParamedicId) > 0 {
			where = append(where, fmt.Sprintf("paramedic_id = $%d", iterate))
			params = append(params, request.ParamedicId)
			iterate = iterate + 1
		}
		if len(request.PatientId) > 0 {
			where = append(where, fmt.Sprintf("patient_id = $%d", iterate))
			params = append(params, request.PatientId)
		}
	}

	finalQuery := qry
	if len(where) > 0 {
		finalQuery = finalQuery + " WHERE " + strings.Join(where, " AND ")
	}
	finalQuery = finalQuery + " ORDER BY create_at DESC"
	log.Printf("LOG FINAL QUERY := %s", finalQuery)
	err := db.Select(&appointments, finalQuery, params...)
	if err != nil {
		log.Printf("ERR SQL = %+v", err)
	}

	return &appointments, nil
}
