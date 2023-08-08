package main

import (
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	appointments "github.com/kalpawreska/jec-appointment/domain/appointment"
)

func main() {
	dtAppointment := time.Now().AddDate(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	dtDateTimeNow := time.Now()

	dtAppointmentProto, errApp := ptypes.TimestampProto(dtAppointment)
	dtNowProto, err := ptypes.TimestampProto(dtDateTimeNow)

	if err != nil {
		log.Println(errApp)
	}

	if err != nil {
		log.Println(err)
	}

	var oAppointment = appointments.AppointmentProto{
		HealthcareId:    "002",
		AppointmentNo:   "AP230807-00001",
		ParamedicId:     "J0001",
		PatientId:       "KP202309-0001",
		AppointmentDate: dtAppointmentProto,
		AppointmentTime: dtNowProto,
		IsVoid:          false,
		UserCreate:      "21239",
		CreateAt:        dtNowProto,
		ScheduleSlotId:  1,
	}

	var oAppointmentList = appointments.AppointmentListProto{Appointments: []*appointments.AppointmentProto{&oAppointment}}

	log.Printf("data: \n%v", oAppointmentList.Appointments)
}
