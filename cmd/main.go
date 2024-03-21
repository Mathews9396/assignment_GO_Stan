package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
	"github.com/Mathews9396/go-userMngmnt/pkg/routes"
)

func main()  {
	router := gin.Default()

	routes.UserManagementRoutes(router)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}