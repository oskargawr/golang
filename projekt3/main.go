package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"projekt3/handlers"
	"projekt3/models"
)

func postPosts(posts []models.Post) {
	for _, post := range posts {
		postWithoutID := models.PostWithoutID{
			Date:     post.Date,
			Country:  post.Country,
			Area:     post.Area,
			Activity: post.Activity,
			Injury:   post.Injury,
		}

		postJSON, err := json.Marshal(postWithoutID)
		if err != nil {
			log.Fatalf("Error marshalling post: %v", err)
		}

		resp, err := http.Post("http://localhost:8080/sharks", "application/json", bytes.NewBuffer(postJSON))
		if err != nil {
			log.Fatalf("Error sending request: %v", err)
		}

		if resp.StatusCode != http.StatusCreated {
			body, _ := ioutil.ReadAll(resp.Body)
			log.Fatalf("Handler returned wrong status code: got %v want %v, body: %s", resp.StatusCode, http.StatusCreated, body)
		}
	}
}

func main() {
	// Load 10 random posts from the JSON file
	posts, err := LoadRandomPosts("data/global-shark-attack.json", 10)
	if err != nil {
		log.Fatalf("Error loading posts: %v", err)
	}
	// print the posts
	for _, post := range posts {
		postJSON, err := json.MarshalIndent(post, "", "  ")
		if err != nil {
			log.Fatalf("Error marshalling post: %v", err)
		}
		fmt.Println(string(postJSON))
	}

	// Dodaj obsługę endpointów do serwera
	http.HandleFunc("/sharks", handlers.PostsHandler)
	http.HandleFunc("/sharks/", handlers.PostHandler)

	fmt.Println("Server is listening on port 8080")

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()
	postPosts(posts)
	select {}
}
