package database

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"mocae/internal/comments"
	"mocae/internal/events"
	"mocae/internal/events_status"
	"mocae/internal/hosts"
	"mocae/internal/logger"
	"mocae/internal/statuses"
	"mocae/internal/users"
)

var DBInstance *gorm.DB
var DBError error

func Init() {
	fmt.Println("Init DB")
	CheckDbUpdate()
}

func DatabaseConnect() {
	logger.LogMsg("Try to connect to DB", "info")
	DBInstance, DBError = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if DBError != nil {
		logger.LogMsg(fmt.Sprintf("Error to Connecting to DB : %s", DBError.Error()), "fatal")
		//fmt.Println(DBError.Error())
		panic("failed to connect database")
	}
	logger.LogMsg("DB Connection is successful", "info")
	//fmt.Println(DBInstance)

}

func CheckDbUpdate() {

	// Migrate the schema
	DBInstance.AutoMigrate(
		&users.Users{}, &hosts.Host{}, &events.Event{}, &comments.Comment{}, &events_status.EventStatus{},
		&statuses.Status{},
	)
	//DBInstance.Migrator().DropColumn(&hosts.Host{}, "comments")

}
