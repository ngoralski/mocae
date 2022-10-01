package users

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"mocae/internal/logger"
	"net/http"
)

//var apiResult common.ApiResult

//var ApiResult struct{}

type Users struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Password  string
	Active    bool
	Apikey    string
	LastLogin string
}

type apiResult struct {
	Message      string  `json:"Message"`
	Status       string  `json:"Status"`
	ReturnedRows int64   `json:"ReturnedRows"`
	DeletedRows  int64   `json:"DeletedRows"`
	UpdatedRows  int64   `json:"UpdatedRows"`
	InsertedRows int64   `json:"InsertedRows"`
	Data         []Users `json:"Data"`
	//Data         []struct{} `json:"Data"`

}

var Instance *gorm.DB
var DBError error

func AllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []Users
	var apiResult apiResult

	result := Instance.Find(&users)
	apiResult.Data = users
	apiResult.Message = "fetch records"
	apiResult.Status = "ok"
	apiResult.ReturnedRows = result.RowsAffected
	json.NewEncoder(w).Encode(apiResult)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user Users
	var apiResult apiResult

	_ = json.NewDecoder(r.Body).Decode(&user)

	result := Instance.Find(&user, "username = ?", user.Username)

	if result.RowsAffected == 0 {
		Instance.Create(&user)
		if Instance.Error != nil {

			logger.LogMsg(fmt.Sprintf("SQL Error : Insert : ", Instance.Error), "info")
			w.WriteHeader(http.StatusFailedDependency)
			apiResult.Message = "Error on inserting data"
			apiResult.Status = "error"

		} else {
			w.WriteHeader(http.StatusCreated)
			apiResult.Data[0] = user
			apiResult.Message = "fetch records"
			apiResult.Status = "ok"
			apiResult.InsertedRows = 1

			json.NewEncoder(w).Encode(apiResult)

		}

	} else {
		logger.LogMsg(fmt.Sprintf("Can't create User %s, it already exist", user.Username), "info")
		w.WriteHeader(http.StatusConflict)

		apiResult.Message = "User already exist"
		apiResult.Status = "error"

		json.NewEncoder(w).Encode(apiResult)

	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user Users
	var apiResult apiResult

	vars := mux.Vars(r)

	result := Instance.Find(&user, "username = ?", vars["username"])

	if result.RowsAffected == 1 {
		user.Active = false
		Instance.Save(&user)

		result := Instance.Delete(&user)

		logger.LogMsg(fmt.Sprintf("Sucessfully deleted User %s", user.Username), "info")
		w.WriteHeader(http.StatusAccepted)

		apiResult.Message = "User deleted"
		apiResult.Status = "ok"
		apiResult.DeletedRows = result.RowsAffected

		json.NewEncoder(w).Encode(apiResult)
	} else {
		logger.LogMsg(fmt.Sprintf("User to deleted %s not defined in db", user.Username), "info")
		w.WriteHeader(http.StatusExpectationFailed)

		apiResult.Message = "Can't deleted an unexistant user"
		apiResult.Status = "error"

		json.NewEncoder(w).Encode(apiResult)
	}
}
