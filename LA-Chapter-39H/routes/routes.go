package routes

import (
	"golang-chapter-39/LA-Chapter-39H/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(controller *controller.UserController) *gin.Engine {
	r := gin.Default()

	r.POST("/users", controller.CreateUser)
	// Implement other routes: GET, PUT, DELETE

	return r
}
