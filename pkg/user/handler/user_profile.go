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

func ViewProfileHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		log.Printf("Error getting ID form context")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"error":  "id not in context",
		})
		return
	}
	UserID, ok := id.(uint64)
	if !ok {
		log.Printf("error parsing id")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "id parsing issue",
		})
		return
	}

	response, err := client.UserProfile(ctx, &userpb.UserID{Id: uint32(UserID)})
	if err != nil {
		log.Printf("Error getting profile of user %v error: %v", UserID, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v profile fetched successfully", UserID),
		"data":    response,
	})
}

func UpdateProfileHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		log.Printf("error getting id from context")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Id not in context",
		})
		return
	}

	userID, ok := id.(uint64)
	if !ok {
		log.Printf("Error parsing id")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Id parsing issue",
		})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding JSON %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.UserProfileUpdate(ctx, &userpb.ProfileUpdate{
		Userid:    uint32(userID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Username:  user.UserName,
		Dob:       user.DoB,
		Gender:    user.Gender,
		Phone:     user.Phone,
	})

	if err != nil {
		log.Printf("error updating user %v error: %v", user.UserName, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v user updated successfullt", user.UserName),
		"data":    response,
	})
}
