package srvhttp

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"mocae/internal/users"
	"net/http"
)

var Instance *gorm.DB
var DBError error

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	users.Instance = Instance
	users.DBError = DBError

	myRouter.HandleFunc("/users", users.AllUsers).Methods("GET")

	myRouter.HandleFunc("/user/{Username}", users.NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{username}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{username}/{email}", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
