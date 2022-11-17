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
	Roles        []*Role        `gorm:"many2many:user_roles;"`
	Applications []*Application `gorm:"foreignKey:LeaderID"`
}

type Role struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Description sql.NullString
	Users       []*User `gorm:"many2many:user_roles;"`
}
