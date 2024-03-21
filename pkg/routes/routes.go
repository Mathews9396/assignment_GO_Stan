package routes

import (
	"github.com/Mathews9396/go-userMngmnt/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func UserManagementRoutes(router *gin.Engine) {
    // Create a group for /book routes
    userRoutes := router.Group("/users")

    // Define routes and corresponding handler functions
    userRoutes.POST("/create", controllers.CreateUser)
    userRoutes.GET("/getAll", controllers.GetAllUsers)
    userRoutes.GET("/find/:userId", controllers.GetUser)
    userRoutes.PUT("/update/:userId", controllers.UpdateUser)
    userRoutes.DELETE("/delete/:userId", controllers.DeleteUser)
}