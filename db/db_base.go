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
		log.Printf("%s %s %d %d\n", item.Username, item.Password, item.ID, item.Type)
	}
	return item
}



type DoctorFilter struct {
	Name string
	Department string
	Sex string
	Age []int
	Page int
}

const PageMax = 10

func GetDoctorsByFilter(filter DoctorFilter) []Doctor {
	var doctors []Doctor
	DB.Table("doctors").
		Where("department like ? AND sex = ?", "%"+filter.Department+"%", filter.Sex).
		Where("age >= ? AND age <= ?", filter.Age[0], filter.Age[1]).
		Limit(PageMax).
		Offset((filter.Page-1)*PageMax).
		Find(&doctors)

	return doctors
}

func GetDoctorsCountByFilter(filter DoctorFilter) int {
	var count int
	DB.Table("doctors").
		Where("department like ? AND sex = ?", "%"+filter.Department+"%", filter.Sex).
		Where("age >= ? AND age <= ?", filter.Age[0], filter.Age[1]).
		Count(&count)
	return count
}

func GetDoctorByID(ID int) Doctor {
	doctor := Doctor{ID: 0}
	DB.Table("doctors").
		Where("id = ?", ID).
		First(&doctor)

	if doctor.ID == 0 {
		log.Warn("get doctor by id from DB fail")
	}
	return doctor
}

func GetDoctorsTotalPageByFilter(filter DoctorFilter) int {
	count := GetDoctorsCountByFilter(filter)
	if count == 0 {
		return 1
	}

	if count%PageMax == 0 {
		return count/PageMax
	} else {
		return count/PageMax + 1
	}
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
