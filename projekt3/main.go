package main

import (
	"fmt"
	"log"
	"net/http"
	"projekt3/handlers"
)

func main() {
	// Load 10 random posts from the JSON file
	posts, err := LoadRandomPosts("data/global-shark-attack.json", 10)
	if err != nil {
		log.Fatalf("Error loading posts: %v", err)
	}

	// Add the posts to the global posts map
	for _, post := range posts {
		handlers.AddPost(post)
	}

	http.HandleFunc("/sharks", handlers.PostsHandler)
	http.HandleFunc("/sharks/", handlers.PostHandler)

	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
