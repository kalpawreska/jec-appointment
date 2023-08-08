package config

var Environment = map[string]interface{}{
	"app_name":    "JEC Appointment gRPC",
	"environment": "production",
	"port":        50051,
	"db_user":     "postgres",
	"db_pass":     "m@nus14sup3r",
	"db_host":     "localhost",
	"db_port":     5432,
	"db_name":     "hacktiv8_go",

	"smtp_host":     "",
	"smtp_port":     "",
	"smtp_from":     "",
	"smtp_username": "",
	"smtp_password": "",
}
