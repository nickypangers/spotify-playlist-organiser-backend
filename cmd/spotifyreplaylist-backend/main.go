package main

import (
	"log"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/server"
)

func main() {
	log.Println("hello world")

	server.InitRouter()

}
