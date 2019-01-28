package main

import (
	"github.com/JohnnyWester/Grabber/news"
	myRouter "github.com/JohnnyWester/Grabber/router"
)

const ServerPost = ":8080"

func main() {
	router := myRouter.New()
	archive := news.New()

	go archive.Serve()
	_ = router.Run(ServerPost)
}
