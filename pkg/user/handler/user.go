package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/godoc-API/pkg/models"
	userpb "github.com/shivaraj-shanthaiah/godoc-API/pkg/user/pb"
)

func UserSignupHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	var user models.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		log.Printf("Eror binding with json :%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	response, err := client.UserSignup(ctx, &userpb.SignupRequest{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Dob:       user.DoB,
		Phone:     user.Phone,
		Gender:    user.Gender,
		Email:     user.Email,
		Password:  user.Password,
	})

	if err != nil {
		log.Printf("error signing up user %v error: %v", user.UserName, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v created successfully", user.UserName),
		"data":    response,
	})
}

func UserLoginHandler(c *gin.Context, client userpb.UserServiceClient, role string) {
	ctx, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	var user models.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding %v", user)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	fmt.Println(user)

	response, err := client.UserLogin(ctx, &userpb.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		log.Printf("Error logging in  user %v error: %v", user.Email, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v logged in successfully", user.Email),
		"data":    response,
	})
}
