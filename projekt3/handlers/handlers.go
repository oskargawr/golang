package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"projekt3/models"
	"sort"
	"strconv"
	"sync"
)

var (
	posts   = make(map[int]models.Post)
	postsMu sync.Mutex
	nextID  = 1
)

func AddPost(p models.Post) {
	postsMu.Lock()
	defer postsMu.Unlock()

	p.ID = new(int)
	*p.ID = nextID
	nextID++
	posts[*p.ID] = p
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetPosts(w, r)
	case "POST":
		handlePostPosts(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/sharks/"):])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		handleGetPost(w, r, id)
	case "PUT":
		handlePutPost(w, r, id)
	case "DELETE":
		handleDeletePost(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetPosts(w http.ResponseWriter, r *http.Request) {
	postsMu.Lock()
	defer postsMu.Unlock()

	ps := make([]models.Post, 0, len(posts))
	for _, p := range posts {
		ps = append(ps, p)
	}

	sort.Slice(ps, func(i, j int) bool {
		return *ps[i].ID < *ps[j].ID
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ps)
}

func handlePostPosts(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var bodyMap map[string]interface{}
	if err := json.Unmarshal(body, &bodyMap); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	// Check if the fields are present in the JSON
	requiredFields := []string{"date", "country", "area", "activity", "injury"}
	for _, field := range requiredFields {
		if _, ok := bodyMap[field]; !ok {
			http.Error(w, "Invalid body structure", http.StatusBadRequest)
			return
		}
	}

	// Check if there are any additional fields
	for field := range bodyMap {
		isRequiredField := false
		for _, requiredField := range requiredFields {
			if field == requiredField {
				isRequiredField = true
				break
			}
		}
		if !isRequiredField {
			http.Error(w, "Invalid body structure", http.StatusBadRequest)
			return
		}
	}

	var p models.Post
	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	AddPost(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Shark created successfully")
}

func handleGetPost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	p, ok := posts[id]
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func handlePutPost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	p, ok := posts[id]
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var bodyMap map[string]interface{}
	if err := json.Unmarshal(body, &bodyMap); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	// Check if the fields are present in the JSON
	requiredFields := []string{"date", "country", "area", "activity", "injury"}
	for _, field := range requiredFields {
		if _, ok := bodyMap[field]; !ok {
			http.Error(w, "Invalid body structure", http.StatusBadRequest)
			return
		}
	}

	// Check if there are any additional fields
	for field := range bodyMap {
		isRequiredField := false
		for _, requiredField := range requiredFields {
			if field == requiredField {
				isRequiredField = true
				break
			}
		}
		if !isRequiredField {
			http.Error(w, "Invalid body structure", http.StatusBadRequest)
			return
		}
	}

	var newP models.Post
	if err := json.Unmarshal(body, &newP); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	newP.ID = p.ID
	posts[id] = newP

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Shark updated successfully"}
	json.NewEncoder(w).Encode(response)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	_, ok := posts[id]
	if !ok {
		http.Error(w, "Shark not found", http.StatusNotFound)
		return
	}

	delete(posts, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": fmt.Sprintf("Successfully deleted shark id %d", id)}
	json.NewEncoder(w).Encode(response)
}
