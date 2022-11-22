package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fbriansyah/app-management/server/model"
	"github.com/fbriansyah/app-management/server/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func runServer(db *gorm.DB) {
	r := gin.Default()

	route.Routing(r, db)

	r.Run()
}

func migrate(mode string, db *gorm.DB, env map[string]string) {
	// function call migration script (create/delete table).
	switch mode {
	case "init":
		// create assosiative table with migration
		model.Init(db)
	case "down":
		model.Purge(db, env["DB_DATABASE"])
	default:
		log.Fatalln("\nUnknown mode. Mode available: \n\t- init (init database) \n\t- down (delete data in database)")
	}
}

func initDB(env map[string]string) *gorm.DB {
	host := env["DB_HOST"]
	port := env["DB_PORT"]
	user := env["DB_USERNAME"]
	pass := env["DB_PASSWORD"]
	dbname := env["DB_DATABASE"]

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		dbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func main() {

	// load .env data
	var envData map[string]string
	envData, err := godotenv.Read()
	if err != nil {
		panic(err)
	}

	// init database
	db := initDB(envData)

	// set default command
	command := ""
	if len(os.Args) > 1 {
		command = os.Args[1]
	} else {
		command = "serve"
	}

	switch command {
	case "migrate":
		if len(os.Args) > 2 {
			migrate(os.Args[2], db, envData)
		} else {
			migrate("up", db, envData)
		}
	case "serve":
		fallthrough
	case "":
		runServer(db)
	default:
		fmt.Println("Unknown command: " + command)
	}
}
