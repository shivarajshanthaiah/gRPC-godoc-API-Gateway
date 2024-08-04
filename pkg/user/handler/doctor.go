package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	userpb "github.com/shivaraj-shanthaiah/godoc-API/pkg/user/pb"
)

func FindAllDoctorsHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	response, err := client.UserFetchAllDoctors(ctxt, &userpb.UNoParam{})
	if err != nil {
		log.Printf("error finding all doctors  err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "fetched all doctors successfully",
		"data":    response,
	})
}

func FindDoctorHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	id := c.Query("id")
	name := c.Query("name")
	response := &userpb.UDoctorModel{}
	var err error
	if id == "" && name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid query",
		})
		return
	} else if id != "" {
		bookID, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "invalid id",
			})
			return
		}
		response, err = client.UserFetchDoctorByID(ctxt, &userpb.UDoctorID{Id: uint32(bookID)})
		if err != nil {
			log.Printf("error finding  doctor by id err: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
	} else if name != "" {
		response, err = client.UserFetchDoctorByName(ctxt, &userpb.UDoctorName{Name: name})
		if err != nil {
			log.Printf("error finding  doctor by id err: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "fetched all doctors successfully",
		"data":    response,
	})

}
