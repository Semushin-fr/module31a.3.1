package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/rss"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/postgres"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type server struct {
	db  storage.Interface
	api *api.API
}

type config struct {
	URLS   []string `json:"rss"`
	Period int      `json:"request_period"`
}

func main() {
	var srv server

	db, err := postgres.New("postgres://postgres:2284@localhost:5432/posts?sslmode=prefer")
	if err != nil {
		log.Fatal(err)
	}

	srv.db = db

	srv.api = api.NewApi(srv.db)

	b, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config config
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatal(err)
	}

	chPosts := make(chan []storage.Post)
	chErrs := make(chan error)
	for _, url := range config.URLS {
		go parseURL(url, chPosts, chErrs, config.Period)
	}

	go func() {
		for posts := range chPosts {
			srv.db.StorePosts(posts)
		}
	}()

	go func() {
		for err := range chErrs {
			log.Println("ошибка:", err)
		}
	}()

	http.ListenAndServe(":8080", srv.api.Router())
}

func parseURL(url string, posts chan<- []storage.Post, errs chan<- error, period int) {
	for {
		news, err := rss.Parse(url)
		if err != nil {
			errs <- err
			continue
		}
		posts <- news
		time.Sleep(time.Minute * time.Duration(period))
	}
}
