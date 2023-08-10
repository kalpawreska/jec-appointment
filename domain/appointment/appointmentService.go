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
	AddRepo(ctx context.Context, app Appointment) (p_strAppointmentNo string, err error)
	GetSingleRepo(ctx context.Context, appointmentNo string, healthcareID string) (*Appointment, error)
	ListRepo(ctx context.Context, request *AppointmentRequest) (*[]Appointment, error)
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

func (asvc appointmentService) GetSingleService(ctx context.Context, appointmentID string, healthcareID string) (*Appointment, error) {
	appointment, err := asvc.repo.GetSingleRepo(ctx, appointmentID, healthcareID)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (asvc appointmentService) ListService(ctx context.Context, request *AppointmentRequest) (*[]Appointment, error) {
	appointments, err := asvc.repo.ListRepo(ctx, request)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

// #endregion end of Business Logic Method
