package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"GoNews/pkg/storage/mongo"
	"GoNews/pkg/storage/postgres"
	"log"
	"net/http"
)

type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	var srv server

	db := memdb.NewStore()

	db2, err := postgres.New("postgres://postgres@localhost:5432/news?sslmode=prefer")
	if err != nil {
		log.Fatal(err)
	}

	db3, err := mongo.New("mongodb://localhost:27017", "news")
	if err != nil {
		log.Fatal(err)
	}

	_, _, _ = db, db2, db3

	srv.db = db2

	srv.api = api.NewApi(srv.db)

	http.ListenAndServe(":8080", srv.api.Router())
}
