package model

import (
	"database/sql"
	"fmt"
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

func Purge(db *gorm.DB, dbName string) {
	queryDrop := fmt.Sprintf("DROP DATABASE `%s`", dbName)
	queryCreate := fmt.Sprintf("CREATE SCHEMA `%s`", dbName)
	// db.Exec(queryDrop)

	db.Connection(func(tx *gorm.DB) error {
		tx.Exec(queryDrop)
		tx.Exec(queryCreate)

		return nil
	})
}
