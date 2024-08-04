package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/godoc-API/pkg/admin/handler"
	adminpb "github.com/shivaraj-shanthaiah/godoc-API/pkg/admin/pb"
)

type AdminRoutes struct {
	client adminpb.AdminServiceClient
}

func (a *AdminRoutes) Login(c *gin.Context) {
	handler.AdminLoginHandler(c, a.client)
}

func (a *AdminRoutes) EditUser(c *gin.Context) {
	handler.EditUserHandler(c, a.client)
}

func (a *AdminRoutes) DeleteUser(c *gin.Context) {
	handler.DeleteUserHandler(c, a.client)
}

func (a *AdminRoutes) FindAllUsers(c *gin.Context) {
	handler.FindAllUsersHandler(c, a.client)
}

func (a *AdminRoutes) FindUserByEmail(c *gin.Context) {
	handler.FindUserByEmailHandler(c, a.client)
}

func (a *AdminRoutes) FindUserByID(c *gin.Context) {
	handler.FindUserByIDHandler(c, a.client)
}

func (a *AdminRoutes) CreateDoctor(c *gin.Context) {
	handler.CreateDoctorHandler(c, a.client)
}

func (a *AdminRoutes) FindAllDoctors(c *gin.Context) {
	handler.FindAllDoctorsHandler(c, a.client)
}

func (a *AdminRoutes) FindDoctor(c *gin.Context) {
	handler.FindDoctorHandler(c, a.client)
}
