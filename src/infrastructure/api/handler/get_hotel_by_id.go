package handler

import (
	"fauna-integration-service/src/infrastructure/fauna"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHotelById(movieRepo fauna.HotelRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()

		println("GET hotel by id")

		hotelId := c.Param("hotelId")

		ret := movieRepo.GetHotelById(hotelId)

		c.JSON(http.StatusOK, ret)
	}
}
