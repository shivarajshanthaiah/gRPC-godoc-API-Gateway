package admin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/godoc-API/middleware"
)

func NewAdminRoute(c *gin.Engine) {
	client, err := ClientDial()
	if err != nil {
		log.Fatalf("Error Not connected with admin gRPC Server, %v", err.Error())
	}

	adminHandler := AdminRoutes{
		client: client,
	}

	apiAdmin := c.Group("/admin")
	{
		apiAdmin.POST("/login", adminHandler.Login)

	}

	apiAdminAuth := c.Group("/admin/auth")
	apiAdminAuth.Use(middleware.Authorization("admin"))
	{
		apiAdminAuth.PUT("/edit/user", adminHandler.EditUser)
		apiAdminAuth.DELETE("delete/user/:id", adminHandler.DeleteUser)
		apiAdminAuth.GET("/users", adminHandler.FindAllUsers)
		apiAdminAuth.GET("/:id/user", adminHandler.FindUserByID)
		apiAdminAuth.GET("/email/user", adminHandler.FindUserByEmail)
		apiAdminAuth.POST("/add/doctor", adminHandler.CreateDoctor)
		apiAdminAuth.GET("/doctors", adminHandler.FindAllDoctors)
		apiAdminAuth.GET("/doctor", adminHandler.FindDoctor)
	}
}
