package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vijishetty29/Go/prepim-amap-api/controllers"
	"github.com/vijishetty29/Go/prepim-amap-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	fmt.Println("Started Application!")
	r := gin.Default()
	r.GET("/amap-services-go/selections", controllers.GetSelectionsData)
	r.Run("localhost:3000")
}
