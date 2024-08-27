package routes

import (
	"golang-chapter-39/LA-Chapter-39H-I/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	r.POST("/users", ctx.Ctl.User.CreateUser)
	r.GET("/users/:id", ctx.Ctl.User.GetUser)
	r.PUT("/users/:id", ctx.Ctl.User.UpdateUser)
	r.DELETE("/users/:id", ctx.Ctl.User.DeleteUser)

	return r
}
