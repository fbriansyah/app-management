package model

import (
	"database/sql"
	"log"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&Role{},
		&Application{},
		&Server{},
		&Cluster{},
	)

	err := db.SetupJoinTable(&Application{}, "Servers", &Cluster{})
	if err != nil {
		log.Fatalln("Application", err)
	}
	err = db.SetupJoinTable(&Server{}, "Apps", &Cluster{})
	if err != nil {
		log.Fatalln("Server", err)
	}

	// create default role
	roles := []Role{
		{Name: "User", Description: sql.NullString{String: "Default User", Valid: true}},
		{Name: "Admin", Description: sql.NullString{String: "Administration User", Valid: true}},
	}
	db.Create(&roles)
}
