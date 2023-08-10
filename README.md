# JEC Appointment
JEC Appointment is simple Go Language Programming Backend for patient appointment.

# Service List :
1. Create Appointment 		: POST {BASE_URL_API}/api/appointment [v0.0.2]
2. Read All Appointment 	: GET {BASE_URL_API}/api/appointment [v0.0.1]
3. Read Appointment 		: GET {BASE_URL_API}/api/appointment/{appointment_id} [v0.0.1]


### Run Swagger
```
swag init -pd -g "./cmd/api/main.go"
```

### Run App
```
go run .\cmd\api\