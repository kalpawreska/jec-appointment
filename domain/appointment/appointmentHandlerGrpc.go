// Declares the package name
package appointment

//	Import library
import (
	context "context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// #region Trademark

// This software, all associated documentation, and all copies are CONFIDENTIAL INFORMATION of Kalpavriksha
// https://www.fwahyudianto.id
// ® Wahyudianto, Fajar
// Email 	: me@fwahyudianto.id

// #endregion

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
	dataResult, err := gh.appointmentSvc.ListService(ctx)
	if len(dataResult) == 0 {
		return
	}

	var resultdata []*AppointmentProto
	for _, v := range dataResult {
		resultdata = append(resultdata, &AppointmentProto{
			HealthcareId:    v.HealthcareId,
			AppointmentNo:   v.AppointmentNo,
			ParamedicId:     v.ParamedicId,
			PatientId:       v.PatientId,
			AppointmentDate: timestamppb.New(v.AppointmentDate),
			AppointmentTime: timestamppb.New(v.AppointmentTime),
			IsVoid:          v.IsVoid,
			UserCreate:      v.UserCreate,
			CreateAt:        timestamppb.New(v.CreateAt),
			ScheduleSlotId:  int64(v.ScheduleSlotId),
		})
	}
	res = &AppointmentListProto{Appointments: resultdata}

	return
}

func (gh *appointmentGrpcHandler) ListByParam(ctx context.Context, filter *AppointmentProto) (res *AppointmentListProto, err error) {
	// res = new(AppointmentListProto)
	// appointmentRequest := AppointmentRequest{
	// 	HealthcareId:  filter.HealthcareId,
	// 	ParamedicId:   filter.ParamedicId,
	// 	PatientId:     filter.PatientId,
	// 	AppointmentNo: filter.AppointmentNo,
	// }
	// appointments, err := gh.appointmentSvc.ListService(ctx, &appointmentRequest)
	// if err != nil {
	// 	return nil, err
	// }

	// appointmentProtos := []*AppointmentProto{}
	// for _, appt := range *appointments {
	// 	apptProto := appt.ParseToAppointmentProto()
	// 	appointmentProtos = append(appointmentProtos, apptProto)
	// }
	// res.Appointments = appointmentProtos

	return
}

func (gh *appointmentGrpcHandler) Get(ctx context.Context, filter *AppointmentGetProto) (res *AppointmentProto, err error) {
	getData := AppointmentRequest{
		HealthcareId:  filter.GetHealthcareId(),
		AppointmentNo: filter.GetAppointmentNo(),
	}

	dataResult, err := gh.appointmentSvc.GetService(ctx, getData)
	if len(dataResult) == 0 {
		return
	}

	res = &AppointmentProto{
		HealthcareId:    dataResult[0].HealthcareId,
		AppointmentNo:   dataResult[0].AppointmentNo,
		ParamedicId:     dataResult[0].ParamedicId,
		PatientId:       dataResult[0].PatientId,
		AppointmentDate: timestamppb.New(dataResult[0].AppointmentDate),
		AppointmentTime: timestamppb.New(dataResult[0].AppointmentTime),
		IsVoid:          dataResult[0].IsVoid,
		UserCreate:      dataResult[0].UserCreate,
		CreateAt:        timestamppb.New(dataResult[0].CreateAt),
	}

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
