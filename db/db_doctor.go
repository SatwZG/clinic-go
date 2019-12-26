package clinicDB

import (
	log "github.com/sirupsen/logrus"
)

type DoctorFilter struct {
	Name string
	Department string
	Sex string
	Age []int
	Page int
}

const PageMax = 10

// 有 BUG
func GetDoctorsByFilter(filter DoctorFilter) []Doctor {
	var doctors []Doctor
	if len(filter.Age) == 0 {
		filter.Age = []int{0, 1000}
	}

	log.Printf("age: %d %d", filter.Age[0], filter.Age[1])
	DB.Table("doctors").
		Where("name like ?", "%"+filter.Name+"%").
		Where("department like ? AND sex like ?", "%"+filter.Department+"%", "%"+filter.Sex+"%").
		Where("age >= ? AND age <= ?", filter.Age[0], filter.Age[1]).
		Limit(PageMax).
		Offset((filter.Page-1)*PageMax).
		Find(&doctors)
	log.Printf("", len(doctors))

	return doctors
}

func GetDoctorsCountByFilter(filter DoctorFilter) int {
	var count int
	if len(filter.Age) == 0 {
		filter.Age = []int{0, 1000000}
	}

	DB.Table("doctors").
		Where("name like ?", "%"+filter.Name+"%").
		Where("department like ? AND sex like ?", "%"+filter.Department+"%", "%"+filter.Sex+"%").
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


func AddDoctor(doctor Doctor, password string) int {
	ID := AddAccount(Account{Username: doctor.Username, Password: password, Type: 1})
	// user zero temporarily
	if ID == 0 {
		return ID
	}
	doctor.ID = ID
	DB.Table("doctors").Create(&doctor)
	return ID
}

func UpdateDoctor(doctor Doctor, password string) int {
	var dbDoctor Doctor
	DB.Table("doctors").
		Where("username = ?", doctor.Username).
		First(&dbDoctor)
	if dbDoctor.ID == 0 {
		log.Warn("UpdateDoctor fail because table doctor not exit")
		return dbDoctor.ID
	}
	if password != "" {
		if UpdateAccount(doctor.Username, password) == 0 {
			log.Warn("doctor not match with account at UpdateDoctor")
			return 0
		}
	}

	dbDoctor.Name = doctor.Name
	dbDoctor.Sex = doctor.Sex
	dbDoctor.Age = doctor.Age
	dbDoctor.Department = doctor.Department
	if doctor.Introduction.Valid {
		dbDoctor.Introduction = doctor.Introduction
	}
	if doctor.Avatar.Valid {
		dbDoctor.Avatar = doctor.Avatar
	}
	DB.Table("doctors").Save(&dbDoctor)

	return dbDoctor.ID
}

func DeleteDoctor(username string) {
	DeleteAccountByUsername(username)
	DB.Where("username = ?", username).
		Delete(Doctor{})
}