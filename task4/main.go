package main

import (
	"github.com/gin-gonic/gin"
	"task4/db"
	"task4/router"
)

func main() {
	db.DB = db.Initdb()
	r := gin.Default()
	router.RegisterRoutes(r)
	r.Run()
}
