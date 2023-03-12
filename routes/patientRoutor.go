package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func PatientRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/addpatient", controller.Addpatient())
	incomingRoutes.DELETE("/delete/patient", controller.DeletePatient())
	incomingRoutes.GET("/showall/doctors", controller.Get_docter())
	incomingRoutes.GET("/get_doctor_by_city", controller.GetDoctorByLocation())
	incomingRoutes.POST("/bookappointment", controller.BookingAppointment())
	incomingRoutes.DELETE("/cancelAppointment", controller.Cancel_appointment())

}
