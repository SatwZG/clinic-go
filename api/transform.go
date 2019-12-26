package api

import (
	"clinic-go/db"
	"database/sql"
)

func DbAdmin2req(admin clinicDB.Admin) Admin {
	return Admin {
		Username: admin.Username,
		Name: admin.Name,
	}
}

func DbDoctor2req(doctor clinicDB.Doctor) Doctor {
	return Doctor {
		Username: doctor.Username,
		Name: doctor.Name,
		Sex: doctor.Sex,
		Age: doctor.Age,
		Department: doctor.Department,
		Avatar: doctor.Avatar.String,
		Introduction: doctor.Introduction.String,
	}
}

func Req2DbDoctor(doctor Doctor) clinicDB.Doctor {
	return clinicDB.Doctor {
		Username: doctor.Username,
		Name: doctor.Name,
		Sex: doctor.Sex,
		Age: doctor.Age,
		Department: doctor.Department,
		Avatar: sql.NullString{String: doctor.Avatar, Valid: doctor.Avatar != ""},
		Introduction: sql.NullString{String: doctor.Introduction, Valid: doctor.Introduction != ""},
	}
}

func Req2DbDoctorfilter(req DoctorSearchRequest) clinicDB.DoctorFilter {
	return clinicDB.DoctorFilter {
		Name: req.Name,
		Department: req.Department,
		Sex: req.Sex,
		Age: req.Age,
		Page: req.Page,
	}
}

func DbMedicine2Req(medicine clinicDB.Medicine) Medicine {
	return Medicine {
		Name: medicine.Name,
		Count: medicine.Count,
	}
}

func Req2DbMedicine(medicine Medicine) clinicDB.Medicine {
	return clinicDB.Medicine {
		Name: medicine.Name,
		Count: medicine.Count,
	}
}