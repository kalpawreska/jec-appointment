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

// #endregion end of Business Logic Method
