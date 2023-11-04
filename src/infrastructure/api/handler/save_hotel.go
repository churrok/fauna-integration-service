package handler

import (
	"fauna-integration-service/src/domain/entities"
	"fauna-integration-service/src/infrastructure/fauna"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveHotel(hotelRepository fauna.HotelRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()

		println("POST hoteles")

		var hotel *entities.Hotel
		if err := c.ShouldBindJSON(&hotel); err != nil {
			c.JSON(http.StatusBadRequest, "wrong json structure")
			return
		}

		ret := hotelRepository.SaveHotel(*hotel)

		c.JSON(http.StatusCreated, ret)
	}
}
