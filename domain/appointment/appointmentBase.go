// Declares the package name
package appointment

//	Import library
import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func RouterInit(r fiber.Router) {
	svc := NewAppointmentService()
	handler := NewAppointmentHandler(svc)

	appointmentApi := r.Group("/api")
	appointmentApi.Get("/appointment", handler.List)
	appointmentApi.Get("/appointment/get", handler.Get)
	appointmentApi.Post("/appointment", handler.Create)
	appointmentApi.Put("/appointment", handler.Update)
	appointmentApi.Delete("/appointment", handler.Delete)
}

func RouterInitWithDB(r fiber.Router, dbx *sqlx.DB) {
	repo := NewAppointmentRepository(dbx)
	svc := NewAppointmentServiceWithDB(repo)
	handler := NewAppointmentHandler(svc)

	appointmentApi := r.Group("/api")
	appointmentApi.Get("/appointment", handler.List)
	appointmentApi.Get("/appointment/get", handler.Get)
	appointmentApi.Post("/appointment", handler.Create)
	appointmentApi.Put("/appointment", handler.Update)
	appointmentApi.Delete("/appointment", handler.Delete)
}

func RouterInitGRPC(server *grpc.Server, db *sqlx.DB) {
	repo := NewAppointmentRepository(db)
	svc := NewAppointmentServiceWithDB(repo)
	handler := NewAppointmentGrpcHandler(svc)

	RegisterAppointmentServiceServer(server, &handler)
}
