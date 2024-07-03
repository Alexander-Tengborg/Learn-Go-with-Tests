package main

import (
	"log"
	"os"

	blogposts "github.com/alexander-tengborg/learn-go-with-tests/files"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(posts)
}
