package clinicDB

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var DB *gorm.DB

func InitDB() {
	// CLINIC_DB='host=localhost database=clinic  user=postgres password= sslmode=disable'
	log.Println("start init DB")
	myDB, err := gorm.Open("postgres", os.Getenv("CLINIC_DB"))
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



func GetAccountByUsername(username string) Account {
	item := Account{ID: 0}
	DB.Table("accounts").
		Where("username = ?", username).
		First(&item)
	log.Printf("%s %s %d %d\n", item.Username, item.Password, item.ID, item.Type)
	return item
}

func GetAccountByID(ID int) Account {
	item := Account{ID: 0, Type: 0}
	DB.Table("accounts").
		Where("id = ?", ID).
		First(&item)
	if item.ID == 0 {
		log.Warn("GetAccountByID fail, it is impossible")
	} else {
		log.Printf("op: %s %s %d %d\n", item.Username, item.Password, item.ID, item.Type)
	}
	return item
}

func AddAccount(account Account) int {
	DB.Table("accounts").Create(&account)
	if account.ID == 0 {
		log.Warn("", account.ID)
	}
	return account.ID
}

func UpdateAccount(username string, password string) int {
	var account Account
	DB.Table("accounts").
		Where("username = ?", username).
		First(&account)
	if account.ID == 0 {
		log.Warn("", account.ID)
	}
	account.Password = password
	DB.Table("accounts").Save(&account)
	return account.ID
}

func DeleteAccountByUsername(username string) {
	DB.Where("username = ?", username).
		Delete(Account{})
}



func GetAdminByID(ID int) Admin {
	admin := Admin{ID: 0}
	DB.Table("admins").
		Where("id = ?", ID).
		First(&admin)
	if admin.ID == 0 {
		log.Warn("get admin by id from DB fail")
	}
	return admin
}
