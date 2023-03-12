package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func DoctorRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/addDoctor", controller.Add_docter())          //done
	incomingRoutes.PUT("/update/doctor", controller.Update_docter())    //done
	incomingRoutes.DELETE("/delete/doctor", controller.Delete_docter()) //done
}
