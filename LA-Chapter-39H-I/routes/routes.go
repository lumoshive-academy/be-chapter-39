package routes

import (
	"golang-chapter-39/LA-Chapter-39H-I/infra"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	r.POST("/users", ctx.Ctl.User.CreateUser)
	// Implement other routes: GET, PUT, DELETE

	return r
}
