package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"unique"`
	Password     string
	Phone        sql.NullString
	Roles        []*Role        `gorm:"many2many:user_roles"`
	Applications []*Application `gorm:"foreignKey:LeaderID"`
}

type Role struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Users       []*User        `gorm:"many2many:user_roles"`
}
