package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"projekt3/models"
	"time"
)

func LoadRandomPosts(filename string, count int) ([]models.Post, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var allPosts []models.Post
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &allPosts); err != nil {
		return nil, err
	}

	if len(allPosts) < count {
		count = len(allPosts)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allPosts), func(i, j int) { allPosts[i], allPosts[j] = allPosts[j], allPosts[i] })

	return allPosts[:count], nil
}
