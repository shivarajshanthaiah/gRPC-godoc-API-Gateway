package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	adminpb "github.com/shivaraj-shanthaiah/godoc-API/pkg/admin/pb"
	"github.com/shivaraj-shanthaiah/godoc-API/pkg/models"
)

func CreateDoctorHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	var doctor models.Doctor

	if err := c.ShouldBindJSON(&doctor); err != nil {
		log.Printf("error binding json :%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.CreateDoctor(ctxt, &adminpb.DoctorModel{
		DoctorName: doctor.Name,
		Age:        doctor.Age,
		Speciality: doctor.Speciality,
		Gender:     doctor.Gender,
		Email:      doctor.Email,
	})
	if err != nil {
		log.Printf("error creating doctor %v err: %v", doctor.Name, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v created successfully", doctor.Name),
		"data":    response,
	})
}

func FindAllDoctorsHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	response, err := client.FetchAllDoctors(ctxt, &adminpb.AdminNoParam{})
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

func FindDoctorHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	id := c.Query("id")
	name := c.Query("name")
	response := &adminpb.DoctorModel{}
	var err error
	if id == "" && name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid query",
		})
		return
	} else if id != "" {
		doctorID, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "invalid id",
			})
			return
		}
		response, err = client.UserFetchDoctorByID(ctxt, &adminpb.DoctorID{Id: uint32(doctorID)})
		if err != nil {
			log.Printf("error finding  doctor by id err: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
	} else if name != "" {
		response, err = client.UserFetchDoctorByName(ctxt, &adminpb.DoctorName{Name: name})
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
