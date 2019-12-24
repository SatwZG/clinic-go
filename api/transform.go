package api

import (
	"clinic-go/db"
)

func DbAdmin2req(admin clinicDB.Admin) Admin {
	return Admin {
		ID: admin.ID,
		Username: admin.Username,
		Name: admin.Name,
	}
}

func DbDoctor2req(doctor clinicDB.Doctor) Doctor {
	return Doctor {
		ID: doctor.ID,
		Username: doctor.Username,
		Name: doctor.Name,
		Sex: doctor.Sex,
		Age: doctor.Age,
		Department: doctor.Department,
		Avatar: doctor.Avatar.String,
		Introduction: doctor.Introduction.String,
	}
}

func req2DbDoctorfilter(req DoctorSearchRequest) clinicDB.DoctorFilter {
	return clinicDB.DoctorFilter {
		Name: req.Name,
		Department: req.Department,
		Sex: req.Sex,
		Age: req.Age,
		Page: req.Page,
	}
}