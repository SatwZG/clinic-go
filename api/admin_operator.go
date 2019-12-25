package api

import (
	"clinic-go/db"
)



func GetDoctorsBySearch(req DoctorSearchRequest) []clinicDB.Doctor {
	return clinicDB.GetDoctorsByFilter(Req2DbDoctorfilter(req))
}

func GetDoctorsTotalPageBySearch(req DoctorSearchRequest) int {
	return clinicDB.GetDoctorsTotalPageByFilter(Req2DbDoctorfilter(req))
}
