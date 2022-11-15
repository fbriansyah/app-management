package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func runServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func migrate(mode string) {
	// function call migration script (create/delete table).
	switch mode {
	case "up":
		fmt.Println("Migrate Up")
	case "down":
		fmt.Println("Migrate Down")
	default:
		log.Fatalln("\nUnknown mode. Mode available: \n\t- up (init database) \n\t- down (delete data in database)")
	}
}

func main() {
	command := ""

	// load .env data
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// set default command
	if len(os.Args) > 1 {
		command = os.Args[1]
	} else {
		command = "serve"
	}

	switch command {
	case "migrate":
		if len(os.Args) > 2 {
			migrate(os.Args[2])
		} else {
			migrate("up")
		}
	case "serve":
		fallthrough
	default:
		runServer()
	}
}
