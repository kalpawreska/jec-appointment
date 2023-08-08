// Declares the package name
package appointment

//	Import library
import "log"

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

// Declare Appointment Service construct
type appointmentService struct {
}

// Declare Appointment Service
func AppointmentService() appointmentService {
	return appointmentService{}
}

// #region Business Logic Method

// Implement Create Appointment on handler
func (asvc appointmentService) CreateService(req AppointmentRequest) (err error) {
	model := req.ParseToEntity()

	_ = model
	log.Printf("Create Appointment [%v] successfully!\n", model)

	return
}

// #endregion end of Business Logic Method
