package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Store struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Active bool   `json:"ref"`
}

var stores []*Store

func createStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newStore Store

	json.NewDecoder(r.Body).Decode(&newStore)
	newStore.ID = strconv.Itoa(len(stores) + 1)
	stores = append(stores, &newStore)
	json.NewEncoder(w).Encode(newStore)
}

func updateStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	type ParamsProps struct {
		Name string `json:"name"`
	}

	var paramsProps ParamsProps

	err = json.Unmarshal(body, &paramsProps)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	params := mux.Vars(r)

	for _, store := range stores {
		if store.ID == params["id"] {
			store.Name = paramsProps.Name
			output, err := json.Marshal(store)

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			w.Write(output)
			return
		}
	}
}

func deleteStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, value := range stores {
		if value.ID == params["id"] {
			stores[index] = stores[len(stores)-1]
			stores = stores[:len(stores)-1]
			json.NewEncoder(w).Encode(stores)
			return
		}
	}
}

func listStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stores)
}

func showStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range stores {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/stores", createStore).Methods("POST")
	router.HandleFunc("/stores/{id}", updateStore).Methods("PUT")
	router.HandleFunc("/stores/{id}", deleteStore).Methods("DELETE")
	router.HandleFunc("/stores", listStore).Methods("GET")
	router.HandleFunc("/stores/{id}", showStore).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
