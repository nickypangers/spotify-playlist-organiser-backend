package main

import (
	"fmt"

	"github.com/nickypangers/spotifyreplaylist-backend/pkg/server"
)

func main() {
	fmt.Println("hello world")

	server.InitRouter()

}
