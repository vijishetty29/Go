package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vijishetty29/Go/go-sitcon/pkg/controllers"
	"github.com/vijishetty29/Go/go-sitcon/pkg/models"
)

func init() {
	models.ConnectDatabase()
}
func main() {
	fmt.Println("Started Application!")
	r := gin.Default()
	r.GET("/amap-services-go/selections", controllers.GetSelectionsData)
	r.Run("localhost:3000")
}
