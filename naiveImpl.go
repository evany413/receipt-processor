package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
)

var lock sync.Mutex

func RunNaiveImpl() {
	http.HandleFunc("/receipts/process", naiveProcessReceipt)
	http.HandleFunc("/receipts/", naiveGetPoints)
	http.ListenAndServe(":8080", nil)
}

func naiveProcessReceipt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := uuid.NewString()
	points := calculatePoints(receipt)

	lock.Lock()
	db[id] = points
	lock.Unlock()

	response := map[string]string{"id": id}
	json.NewEncoder(w).Encode(response)
}

func naiveGetPoints(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	id, err := getIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lock.Lock()
	points, exists := db[id]
	lock.Unlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	response := map[string]int{"points": points}
	json.NewEncoder(w).Encode(response)
}

func getIDFromPath(path string) (string, error) {
	// Trim the prefix and split the rest by '/'
	trimmed := strings.TrimPrefix(path, "/receipts/")
	parts := strings.Split(trimmed, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("path does not contain an ID and action")
	}
	// The ID should be the first part, and "points" should be the second part
	if parts[1] != "points" {
		return "", fmt.Errorf("invalid action: expected 'points', got %s", parts[1])
	}
	return parts[0], nil
}
