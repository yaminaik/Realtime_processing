package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiURL = "https://jsonplaceholder.typicode.com/posts"

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func fetchData() ([]Post, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func processPosts(posts []Post) {
	for _, post := range posts {
		fmt.Printf("Post ID: %d, Title: %s\n", post.ID, post.Title)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		posts, err := fetchData()
		if err != nil {
			log.Println("Error fetching data:", err)
			http.Error(w, "Error fetching data", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(posts)
	})

	log.Println("Server started on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
