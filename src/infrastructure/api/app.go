package api

import (
	"fauna-integration-service/src/infrastructure/api/handler"
	"fauna-integration-service/src/infrastructure/fauna"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Run(hotelRepo fauna.HotelRepository) error {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	api := gin.Default()

	api.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	api.GET("/hotels/:hotelId", handler.GetHotelById(hotelRepo))
	api.POST("/hotels", handler.SaveHotel(hotelRepo))

	println("starting http server on port: 8080")
	return api.Run(":8080")
}
