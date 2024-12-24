package routes

import (
	"go-practice-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default() // inttialize the routes
	r.GET("/users", controllers.GetUsers)
	r.POST("/createuser", controllers.CreateUser)
	r.GET("/user/:id", controllers.GetUserById)
	r.DELETE("/user/:id", controllers.DeleteById)
	r.GET("/orders", controllers.GetOrders)
	return r

}
