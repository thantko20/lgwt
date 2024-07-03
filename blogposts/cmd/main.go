package main

import (
	"log"
	"os"

	"github.com/thantko20/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFs(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(posts)
}
