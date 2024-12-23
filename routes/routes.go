package routes

import (
	"go-practice-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users", controllers.GetUsers)
	r.POST("/createuser", controllers.CreateUser)
	r.GET("/user/:id", controllers.GetUserById)
	return r

}
