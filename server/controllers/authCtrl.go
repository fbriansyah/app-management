package controllers

import (
	"fmt"
	"net/http"

	"github.com/fbriansyah/app-management/server/model"
	"github.com/fbriansyah/app-management/server/requests"
	"github.com/fbriansyah/app-management/server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthCtrl struct {
	Controller
	DB *gorm.DB
}

func (ac *AuthCtrl) Register(ctx *gin.Context) {
	var req requests.UserAuthRequest

	resp := ac.InitResponse()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.Code = http.StatusBadGateway
		resp.Message = err.Error()
	} else {
		var role model.Role
		ac.DB.Where("id = ?", 1).First(&role)

		user := model.User{
			Name:     "",
			Email:    req.Email,
			Password: "",
		}
		user.Roles = append(user.Roles, &role)
		if result := ac.DB.Create(&user); result.Error != nil {
			resp.Code = http.StatusBadGateway
			resp.Message = result.Error.Error()
		} else {

			salt := []byte(fmt.Sprintf("user_%d", user.ID))
			user.Password = utils.HashFunc(req.Password, salt)
			ac.DB.Save(user)
			resp.Data = user
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (ac *AuthCtrl) Login(ctx *gin.Context) {
	var req requests.UserAuthRequest

	resp := ac.InitResponse()
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.Code = http.StatusBadGateway
		resp.Message = err.Error()
	} else {
		var user model.User
		userRes := ac.DB.
			Preload("Roles").
			Where("email = ?", req.Email).
			First(&user)
		if userRes.Error != nil {
			resp.Code = http.StatusNotFound
			resp.Message = "user tidak terdaftar"
		} else {
			salt := []byte(fmt.Sprintf("user_%d", user.ID))
			if utils.PasswordsMatch(user.Password, req.Password, salt) {
				resp.Data = user
			} else {
				resp.Code = http.StatusNotFound
				resp.Message = "password salah"
			}
		}
	}
	ctx.JSON(http.StatusOK, resp)
}
