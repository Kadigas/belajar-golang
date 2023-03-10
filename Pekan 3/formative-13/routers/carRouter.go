package routers

import (
	"formative-13/controllers"

	"github.com/gin-gonic/gin"
)

func StartSever() *gin.Engine {
	router := gin.Default()

	router.GET("/cars", controllers.IndexCar)

	router.POST("/cars", controllers.CreateCar)

	router.PUT("/cars/:carID", controllers.UpdateCar)

	router.GET("/cars/:carID", controllers.GetCar)

	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
