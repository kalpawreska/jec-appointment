<!-- Package List -->
# go get -u github.com/golang/protobuf/protoc-gen-go
# go get -u github.com/golang/protobuf@v1.5.0
# go get -u google.golang.org/grpc@v1.56.0
# go get -u github.com/jmoiron/sqlx
# go get -u github.com/gofiber/fiber/v2

<!-- Command -->
<!-- ! Generate Proto -->
# protoc --go_out=. ./domain/appointment/proto/*.proto
<!-- ! Generate Proto GRPC -->
# protoc --go_out=. ./domain/appointment/proto/*.proto --go-grpc_out=.
<!-- ! Create File -->
# touch domain/appointment/appointment.go
# touch domain/appointment/appointmentDto.go
# touch domain/appointment/appointmentHandler.go
# touch domain/appointment/appointmentService.go
# touch domain/appointment/appointmentValidator.go
# touch domain/appointment/appointmentBase.go
# touch domain/appointment/appointmentRepository.go
# touch domain/appointment/appointmentHandlerGrpc.go

# touch pkg/jecvalidator/jecvalidator.go
