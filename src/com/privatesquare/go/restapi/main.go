package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// The person Type (more like an object)
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Person does not exists", http.StatusNotFound)
}

func PersonExists(id string) bool{
	isExists := false
	for _, item := range people {
		if item.ID == id {
			isExists = true
		}
	}
	return isExists
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	queryValues := r.URL.Query()
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	if !PersonExists(params["id"]){
		person.ID = params["id"]
		person.Firstname = queryValues.Get("fn")
		person.Lastname = queryValues.Get("ln")
		people = append(people, person)
		json.NewEncoder(w).Encode(people)
	}else {
		http.Error(w, "Person already exists", http.StatusFound)
	}

}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Prince", Lastname: "Link", Address: &Address{City: "Hyrule", State: "Hateno"}})
	people = append(people, Person{ID: "2", Firstname: "Princess", Lastname: "Zelda", Address: &Address{City: "City Z", State: "Castle Hyrule"}})
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}