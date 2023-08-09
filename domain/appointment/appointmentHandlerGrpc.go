// Declares the package name
package appointment

//	Import library
import (
	context "context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Declare Appointment Handler construct
type appointmentGrpcHandler struct {
	// Access Service
	appointmentSvc appointmentService
}

// Declare Appointment GRPC Handler
// @Param appointmentService aka p_oGrpcService
func NewAppointmentGrpcHandler(p_oGrpcService appointmentService) appointmentGrpcHandler {
	return appointmentGrpcHandler{appointmentSvc: p_oGrpcService}
}

func (gh *appointmentGrpcHandler) List(ctx context.Context, empty *emptypb.Empty) (res *AppointmentListProto, err error) {
	return
}

func (gh *appointmentGrpcHandler) ListByParam(ctx context.Context, filter *AppointmentProto) (res *AppointmentListProto, err error) {
	return
}

func (gh *appointmentGrpcHandler) Get(ctx context.Context, filter *AppointmentGetProto) (res *AppointmentProto, err error) {
	return
}

func (gh *appointmentGrpcHandler) Add(ctx context.Context, req *AppointmentAddProto) (res *AppointmentProto, err error) {
	reqData := AppointmentRequest{
		HealthcareId:    req.Addappointment.HealthcareId,
		AppointmentNo:   req.Addappointment.AppointmentNo,
		ParamedicId:     req.Addappointment.ParamedicId,
		PatientId:       req.Addappointment.PatientId,
		AppointmentDate: req.Addappointment.AppointmentDate.AsTime(),
		AppointmentTime: req.Addappointment.AppointmentTime.AsTime(),
		IsVoid:          req.Addappointment.IsVoid,
		UserCreate:      req.Addappointment.UserCreate,
		CreateAt:        req.Addappointment.CreateAt.AsTime(),
		ScheduleSlotId:  int(req.Addappointment.ScheduleSlotId),
	}

	err = gh.appointmentSvc.CreateService(ctx, reqData)
	res = req.Addappointment

	return
}

func (gh *appointmentGrpcHandler) Update(ctx context.Context, req *AppointmentUpdateProto) (res *AppointmentProto, err error) {
	return
}

func (gh *appointmentGrpcHandler) Delete(ctx context.Context, req *AppointmentDeleteProto) (res *AppointmentDeleteResponseProto, err error) {
	return
}

func (gh *appointmentGrpcHandler) mustEmbedUnimplementedAppointmentServiceServer() {}
