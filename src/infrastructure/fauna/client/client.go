package client

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func NewClient(secret string, collectionName string) *f.FaunaClient {
	println("creating db connection")
	client := f.NewFaunaClient(secret, f.Endpoint("https://db.fauna.com"))
	println("testing db connection")
	res, err := client.Query(
		f.Exists(f.Collection(collectionName)))
	if err != nil {
		panic(err)
	}
	println("testing collection existence")
	var exists bool

	if err := res.Get(&exists); err != nil {
		panic(err)
	}
	if !exists {
		println("collection doesnt exist")
		_, err := client.Query(f.CreateDatabase(f.Obj{"name": collectionName}))
		if err != nil {
			panic(err)
		}
	} else {
		println("collection exists")
	}
	return client
}

/*
_ ,err := client.Query(
		f.If(
		f.Exists(f.Database(dbName)),
		false,
		f.CreateDatabase(f.Obj{"name": dbName})))
*/
