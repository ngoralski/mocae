package users

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"mocae/internal/common"
	"mocae/internal/logger"
	"net/http"
)

var Instance *gorm.DB
var DBError error
var apiResult common.ApiResult

//var ApiResult struct{}

type Users struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Password  string
	Active    bool
	Apikey    string
	LastLogin string
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []Users

	result := Instance.Find(&users)
	apiResult.Data = users
	apiResult.Message = "fetch records"
	apiResult.Status = "ok"
	apiResult.ReturnedRows = result.RowsAffected
	json.NewEncoder(w).Encode(apiResult)

}

func NewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user Users

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
