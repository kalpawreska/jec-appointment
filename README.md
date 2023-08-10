# JEC Appointment
JEC Appointment is simple Go Language Programming Backend for patient appointment.

## Features List :
1. Create Appointment 		: POST {BASE_URL_API}/api/appointment [v0.0.1]
2. Read All Appointment 	: GET {BASE_URL_API}/api/appointment [v0.0.2]
3. Read Appointment 		: GET {BASE_URL_API}/api/appointment/{appointment_id} [v0.0.2]

## Environtment Variable

```env
CONFIG_SMTP_HOST="smtp.gmail.com"
CONFIG_SMTP_PORT="587"
CONFIG_AUTH_EMAIL="user@gmail.com"
CONFIG_AUTH_PASSWORD="yoursecretcode"

POSTGRES_HOST="localhost"
POSTGRES_PORT="5432"
POSTGRES_DBNAME="yourdatabase"
POSTGRES_USER="root"
POSTGRES_PASSWORD="yourpassword"
POSTGRES_SSLMODE="disable"

BASE_URL="http://localhost"
BASE_PORT=":9091"
WEB_BASE_PORT=":9090"

GRPC_PORT=":9092"
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/kalpawreska/jec-appointment
```

Update setting in .env file


Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run cmd/grpc/main.go
```
