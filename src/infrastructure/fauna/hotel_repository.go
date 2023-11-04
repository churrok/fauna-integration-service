package fauna

import (
	"fauna-integration-service/src/domain/entities"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

type HotelRepository struct {
	Client *f.FaunaClient
}

func NewHotelRepository(client *f.FaunaClient) HotelRepository {
	return HotelRepository{Client: client}
}

func (r HotelRepository) GetHotelById(id string) entities.Hotel {
	getHotel := f.Get(f.Ref(f.Collection("hotel"), id))

	res, err := r.Client.Query(getHotel)
	if err != nil {
		panic(err)
	}

	hotelRes := toHotel(res, err)

	return hotelRes
}

func (r HotelRepository) SaveHotel(hotel entities.Hotel) entities.Hotel {
	createHotel := f.Create(
		f.Collection("hotel"),
		f.Obj{
			"data": hotel,
		})
	res, err := r.Client.Query(createHotel)
	if err != nil {
		panic(err)
	}

	hotelRes := toHotel(res, err)

	return hotelRes
}

func toHotel(res f.Value, err error) entities.Hotel {
	var hotelRes entities.Hotel
	if err := res.At(f.ObjKey("data")).Get(&hotelRes); err != nil {
		panic(err)
	}
	var id f.RefV
	if err = res.At(f.ObjKey("ref")).Get(&id); err != nil {
		panic(err)
	}
	hotelRes.Id = &id.ID
	return hotelRes
}
