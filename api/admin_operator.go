package api

import (
	"clinic-go/db"
)



func GetDoctorsBySearch(req DoctorSearchRequest) []clinicDB.Doctor {
	return clinicDB.GetDoctorsByFilter(req2DbDoctorfilter(req))
}

func GetDoctorsTotalPageBySearch(req DoctorSearchRequest) int {
	return clinicDB.GetDoctorsTotalPageByFilter(req2DbDoctorfilter(req))
}
