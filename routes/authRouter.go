package routes

import (
	"github.com/HarshitNagpal29/go-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("user/signup", controllers.Signup())
	incomingRoutes.POST("user/login", controllers.Login())
}
