package main

import (
	"encoding/json"
	"net/http"
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

	id := r.URL.Path[len("/receipts/"):]
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
