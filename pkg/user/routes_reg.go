package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/godoc-API/pkg/services"
	"github.com/shivaraj-shanthaiah/godoc-API/pkg/user/handler"
	pb "github.com/shivaraj-shanthaiah/godoc-API/pkg/user/pb"
)

type User struct {
	client pb.UserServiceClient
	Redis  *services.RedisClient
}

func (u *User) Login(c *gin.Context) {
	handler.UserLoginHandler(c, u.client, "user")
}

func (u *User) Signup(c *gin.Context) {
	handler.UserSignupHandler(c, u.client)
}

func (u *User) ViewProfile(c *gin.Context) {
	handler.ViewProfileHandler(c, u.client)
}

func (u *User) UpdateProfile(c *gin.Context) {
	handler.UpdateProfileHandler(c, u.client)
}

func (u *User) FindAllDoctors(c *gin.Context) {
	handler.FindAllDoctorsHandler(c, u.client)
}

func (u *User) FindDoctor(c *gin.Context) {
	handler.FindDoctorHandler(c, u.client)
}
