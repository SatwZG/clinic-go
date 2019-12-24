package clinicDB

import (
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var DB *gorm.DB

func InitDB() {
	// LOGIFIER_DB='host=localhost database=clinic  user=postgres password= sslmode=disable'
	myDB, err := gorm.Open("postgres", os.Getenv("LOGIFIER_DB"))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = myDB.DB().Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	myDB.DB().SetMaxOpenConns(10)
	// myDB.LogMode(true)
	log.Println("Connected to database")
	DB = myDB
}
