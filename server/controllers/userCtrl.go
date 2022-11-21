package controllers

import (
	"net/http"

	"github.com/fbriansyah/app-management/server/model"
	"github.com/fbriansyah/app-management/server/requests"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserCtrl struct {
	Controller
	DB *gorm.DB
}

// Get all user
func (uc *UserCtrl) Get(ctx *gin.Context) {
	resp := uc.InitResponse()

	users := []model.User{}

	res := uc.DB.Find(&users)

	if res.Error != nil {
		resp.Code = http.StatusBadGateway
		resp.Message = res.Error.Error()
	} else {
		resp.Data = users
	}

	ctx.JSON(http.StatusOK, resp)
}

// Store user to database
func (uc *UserCtrl) Store(ctx *gin.Context) {
	var req requests.UserStoreRequest
	resp := uc.InitResponse()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.Code = http.StatusBadGateway
		resp.Message = err.Error()
	} else {
		user := model.User{Name: req.Name, Email: req.Email, Password: req.Password}
		var role model.Role

		roleRes := uc.DB.Where("id = ?", req.RoleID).First(&role)
		if roleRes.Error != nil {
			uc.DB.Where("id = ?", 1).First(&role)
		}

		user.Roles = append(user.Roles, &role)

		if result := uc.DB.Create(&user); result.Error != nil {
			resp.Code = http.StatusBadGateway
			resp.Message = result.Error.Error()
		} else {
			resp.Data = user
		}
	}

	ctx.JSON(http.StatusOK, resp)
}
