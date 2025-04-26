package main

import (
	"gocudoperations/db"
	"gocudoperations/handler"

	"github.com/gin-gonic/gin"

	"fmt"
)

func initializeDB() {
	db.Init()
	fmt.Print("DB Connection completed")
}

func main() {

	initializeDB()

	r := gin.Default()

	r.GET("/getUser", handler.GetUser)
	r.POST("/createUser", handler.CreateUser)
	r.PATCH("/updateUser", handler.UpdateUser)
	r.DELETE("/deleteUser", handler.DeleteUser)

	r.Run(":8090")
}
