// Declares the package name
package appointment

//	Import library
import (
	"context"
	"log"
)

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

// Appointment Repository Interface
type Repository interface {
	ListRepo(ctx context.Context) ([]Appointment, error)
	GetRepo(ctx context.Context, m_strHealthcareId string, m_strAppointmentNo string) ([]Appointment, error)
	AddRepo(ctx context.Context, app Appointment) (p_strAppointmentNo string, err error)
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

// Implement List Appointment on handler
func (asvc appointmentService) ListService(ctx context.Context) ([]AppointmentListResponse, error) {
	list := make([]AppointmentListResponse, 0)

	items, err := asvc.repo.ListRepo(ctx)
	if err != nil {
		return list, err
	}

	list = make([]AppointmentListResponse, len(items))
	for i, item := range items {
		list[i].HealthcareId = item.HealthcareId
		list[i].AppointmentNo = item.AppointmentNo
		list[i].ParamedicId = item.ParamedicId
		list[i].PatientId = item.PatientId
		list[i].AppointmentDate = item.AppointmentDate
		list[i].AppointmentTime = item.AppointmentTime
		list[i].IsVoid = item.IsVoid
		list[i].UserCreate = item.UserCreate
		list[i].CreateAt = item.CreateAt
		list[i].ScheduleSlotId = item.ScheduleSlotId
	}

	return list, nil
}

// Implement Get Appointment on handler
func (asvc appointmentService) GetService(ctx context.Context, req AppointmentRequest) (res []Appointment, err error) {
	model := req.ParseToEntity()

	res, err = asvc.repo.GetRepo(ctx, model.HealthcareId, model.AppointmentNo)
	if err != nil {
		return
	}

	return
}

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

// #endregion end of Business Logic Method
