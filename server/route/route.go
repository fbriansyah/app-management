package route

import (
	"github.com/fbriansyah/app-management/server/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Routing all request handler
func Routing(r *gin.Engine, db *gorm.DB) {
	r.GET("/ver", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ver": "0.1.1",
		})
	})

	apiEnd := "/api"
	apiGrp := r.Group(apiEnd)
	{
		userGrp := apiGrp.Group("/user")
		{
			userCtrl := controllers.UserCtrl{DB: db}

			userGrp.GET("/", userCtrl.Get)
			userGrp.POST("/", userCtrl.Store)
		}

		authGrp := apiGrp.Group("/auth")
		{
			authCtrl := controllers.AuthCtrl{DB: db}

			authGrp.POST("/register", authCtrl.Register)
			authGrp.POST("/login", authCtrl.Login)
		}
	}
}
