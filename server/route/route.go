package route

import (
	"github.com/fbriansyah/app-management/server/model"
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

	r.GET("/test", func(c *gin.Context) {
		// roles := []*model.Role{&model.Role{Name: "Admin"}}

		usr := model.User{
			Name:     "admin",
			Email:    "admin@app.com",
			Password: "1223451",
		}

		roleAdmin := model.Role{}
		db.Where("id = ?", 1).First(&roleAdmin)
		// usr := model.User{}
		// db.Where("email = ?", "febrian@app.com").First(&usr)
		usr.Roles = append(usr.Roles, &roleAdmin)

		db.Create(&usr)
		res := db.Save(&usr)

		if res.RowsAffected > 0 {
			c.JSON(200, gin.H{
				"code":    "00",
				"message": "",
				"data":    usr,
			})
		} else {
			c.JSON(500, gin.H{
				"code":    "01",
				"message": "Error insert",
				"data":    nil,
			})
		}
	})
}
