// Declares the package name
package appointment

//	Import library
import (
	"context"
	"log"
)

// Appointment Repository Interface
type Repository interface {
	AddRepo(ctx context.Context, app Appointment) (p_strAppointmentNo string, err error)
	GetListAppointment(ctx context.Context, limit, page int) ([]Appointment, int, error)
	GetByAppointmentNo(ctx context.Context, appointmentNo, healthCareID string) ([]Appointment, error)
}

// Declare Appointment Service construct
type appointmentService struct {
	// Access Appointment Repository
	repo Repository
}

// Declare Appointment Service
func NewAppointmentService() appointmentService {
	return appointmentService{}
}

func NewAppointmentServiceWithDB(p_oRepo Repository) appointmentService {
	return appointmentService{
		repo: p_oRepo,
	}
}

// #region Business Logic Method

// Implement Create Appointment on handler
func (asvc appointmentService) CreateService(ctx context.Context, req AppointmentRequest) (err error) {
	model := req.ParseToEntity()

	id, err := asvc.repo.AddRepo(ctx, model)
	if err != nil {
		return
	}

	log.Printf("Create Appointment [%v] successfully!\n", id)

	return nil
}

func (asvc appointmentService) ListAppointment(ctx context.Context, limit, page int) ([]AppointmentListResponse, int, error) {
	result := make([]AppointmentListResponse, 0)

	items, total, err := asvc.repo.GetListAppointment(ctx, limit, page)
	if err != nil {
		return result, 0, err
	}

	result = make([]AppointmentListResponse, len(items))
	for i, item := range items {
		result[i].UserCreate = item.UserCreate
		result[i].HealthcareId = item.HealthcareId
		result[i].AppointmentDate = item.AppointmentDate
		result[i].AppointmentTime = item.AppointmentTime
		result[i].ScheduleSlotId = item.ScheduleSlotId
		result[i].IsVoid = item.IsVoid
		result[i].ParamedicId = item.ParamedicId
		result[i].PatientId = item.PatientId
		result[i].CreateAt = item.CreateAt
		result[i].AppointmentNo = item.AppointmentNo
	}

	return result, total, nil
}

func (asvc appointmentService) DetailAppointment(ctx context.Context, appointmentNo, healthCareID string) ([]AppointmentListResponse, error) {
	var result []AppointmentListResponse

	items, err := asvc.repo.GetByAppointmentNo(ctx, appointmentNo, healthCareID)
	if err != nil {
		return result, err
	}

	result = make([]AppointmentListResponse, len(items))
	for i, item := range items {
		result[i].UserCreate = item.UserCreate
		result[i].HealthcareId = item.HealthcareId
		result[i].AppointmentDate = item.AppointmentDate
		result[i].AppointmentTime = item.AppointmentTime
		result[i].ScheduleSlotId = item.ScheduleSlotId
		result[i].IsVoid = item.IsVoid
		result[i].ParamedicId = item.ParamedicId
		result[i].PatientId = item.PatientId
		result[i].CreateAt = item.CreateAt
		result[i].AppointmentNo = item.AppointmentNo
	}
	return result, nil
}

// #endregion end of Business Logic Method
