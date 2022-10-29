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

	// create a serve mux
	sm := mux.NewRouter()
	sm.StrictSlash(true)
	//myRouter := mux.NewRouter().StrictSlash(true)

	// register handlers
	postR := sm.Methods(http.MethodPost).Subrouter()
	getR := sm.Methods(http.MethodGet).Subrouter()
	putR := sm.Methods(http.MethodPut).Subrouter()
	deleteR := sm.Methods(http.MethodDelete).Subrouter()

	users.Instance = Instance
	users.DBError = DBError

	// Define GET Call
	getR.HandleFunc("/users", users.AllUsers)
	//myRouter.HandleFunc("/users", users.AllUsers).Methods("GET")

	// Define POST Call
	postR.HandleFunc("/user/{Username}", users.CreateUser)
	//myRouter.HandleFunc("/user/{Username}", users.CreateUser).Methods("POST")

	// Define DELETE Call
	//myRouter.HandleFunc("/user/{username}", users.DeleteUser).Methods("DELETE")
	deleteR.HandleFunc("/user/{username}", users.DeleteUser)

	// Define PUT Call
	putR.HandleFunc("/user/{username}/{email}", updateUser)
	//myRouter.HandleFunc("/user/{username}/{email}", updateUser).Methods("PUT")

	//// used the PathPrefix as workaround for scenarios where all the
	//// get requests must use the ValidateAccessToken middleware except
	//// the /refresh-token request which has to use ValidateRefreshToken middleware
	//refToken := sm.PathPrefix("/refresh-token").Subrouter()
	//refToken.HandleFunc("", uh.RefreshToken)
	//refToken.Use(uh.MiddlewareValidateRefreshToken)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
