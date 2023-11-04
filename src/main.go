package main

import (
	"fauna-integration-service/src/infrastructure/api"
	"fauna-integration-service/src/infrastructure/fauna"
	"fauna-integration-service/src/infrastructure/fauna/client"
)

func main() {
	faunaClient := client.NewClient("fnAFR8HpAVAAQB4Egd5jaySmMU-JAZ48T2oEVRXZ", "hotel")
	hotelRepo := fauna.NewHotelRepository(faunaClient)

	err := api.Run(hotelRepo)
	if err != nil {
		panic(err)
	}
}
