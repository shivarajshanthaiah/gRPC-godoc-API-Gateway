package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/godoc-API/middleware"
	"github.com/shivaraj-shanthaiah/godoc-API/pkg/services"
)

func NewUserRoute(c *gin.Engine) {
	client, err := ClientDial()
	redisClient := services.NewRedisClient()
	if err != nil {
		log.Fatalf("error not connected with gRPC server, %v", err.Error())
	}

	userHandler := User{
		client: client,
		Redis:  redisClient,
	}
	apiUser := c.Group("/user")
	{
		apiUser.POST("/login", userHandler.Login)
		apiUser.POST("/signup", userHandler.Signup)
	}

	apiUserAuth := c.Group("/user/auth")
	apiUserAuth.Use(middleware.Authorization("user"))
	{
		apiUserAuth.GET("/profile/view", userHandler.ViewProfile)
		apiUserAuth.POST("/profile/update", userHandler.UpdateProfile)
		apiUserAuth.GET("/doctors", userHandler.FindAllDoctors)
		apiUserAuth.GET("/doctor", userHandler.FindDoctor)
	}
}
