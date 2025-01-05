package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reservation-system/util"
)

func AddStore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	var newStore util.Store
	err = json.Unmarshal(body, &newStore)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	newStore.Queue = make(chan string, newStore.Limit)
	util.Stores = append(util.Stores, newStore)
	fmt.Printf("Received new store: %+v\n", newStore)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Store created successfully"))
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	var newUser util.User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
	}

	for i := range util.Stores {
		if util.Stores[i].Id == newUser.StoreId {
			if addStoreChannel(&util.Stores[i], newUser) {
				util.Users = append(util.Users, newUser)
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("User added to queue successfully"))
				return
			} else {
				http.Error(w, "Store's queue is full", http.StatusServiceUnavailable)
				return
			}
		}
	}
	http.Error(w, "Store not found", http.StatusNotFound)
}

func addStoreChannel(store *util.Store, user util.User) bool {
	select {
	case store.Queue <- user.Id:
		fmt.Printf("User %s added to store %s's queue\n", user.Id, store.Id)
		return true
	default:
		fmt.Printf("Store's queue is full\n")
		return false
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	storeId := r.URL.Query().Get("storeId")
	if storeId == "" {
		http.Error(w, "Missing storeId in query parameter", http.StatusBadRequest)
		return
	}
	fmt.Printf("Received storeId: %s\n", storeId)
	var userId string
	for i := range util.Stores {
		if util.Stores[i].Id == storeId {
			userId = getUserFromChannel(&util.Stores[i])
		}
	}
	if userId == "" {
		http.Error(w, "No users in the queue", http.StatusNotFound)
		return
	}

	var name string
	var number string
	for _, user := range util.Users {
		if user.Id == userId {
			name = user.Name
			number = user.MobileNo
		}
	}

	response := util.UserResponse{
		UserId:   userId,
		Name:     name,
		MobileNo: number,
		Message:  "User processed and deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON", http.StatusInternalServerError)
	}
}

func getUserFromChannel(store *util.Store) string {
	var userId string
	select {
	case userId = <-store.Queue:
		fmt.Printf("Retrieved user ID: %s from queue\n", userId)
		return userId
	default:
		return ""
	}
}
