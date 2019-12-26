package api

import (
	"clinic-go/db"
	"database/sql"
	"encoding/json"

	log "github.com/sirupsen/logrus"
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

func Req2DbMedicines(medicine []Medicine) []clinicDB.Medicine {
	dbMedicines := make([]clinicDB.Medicine, len(medicine))
	for i := 0; i < len(medicine); i++ {
		dbMedicines[i] = Req2DbMedicine(medicine[i])
	}
	return dbMedicines
}


// prescription
func Req2DbPrescriptionFilter(req SearchPrescriptionsWithPageRequest) clinicDB.PrescriptionFilter {
	return clinicDB.PrescriptionFilter {
		Department: req.DoctorName,
		DoctorName: req.DoctorName,
		PatientName: req.PatientName,
		Age: req.Age,
		Sex: req.Sex,
		Time: req.Time,
		Page: req.Page,
	}
}

func DbPrescription2Req(prescription clinicDB.Prescription) Prescription {

	reqPrescription := Prescription {
		ID: prescription.ID,
		Department: prescription.Department,
		DoctorName: prescription.DoctorName,
		PatientName: prescription.PatientName,
		Age: prescription.Age,
		Sex: prescription.Sex,
		CreateTime: prescription.CreateTime,
	}

	if prescription.Medicines != "" {
		err := json.Unmarshal([]byte(prescription.Medicines), &reqPrescription.Medicines)
		if err != nil {
			log.Warn("transform json to Medicines fail")
		}
	}

	return reqPrescription
}

func Req2DbPrescription(prescription Prescription) clinicDB.Prescription {
	dbPrescription := clinicDB.Prescription {
		Department: prescription.Department,
		DoctorName: prescription.DoctorName,
		PatientName: prescription.PatientName,
		Age: prescription.Age,
		Sex: prescription.Sex,
		CreateTime: prescription.CreateTime,
	}

	if len(prescription.Medicines) > 0 {
		bytes, err := json.Marshal(prescription.Medicines)
		if err != nil {
			log.Warn("transform Medicines to json fail")
		} else {
			dbPrescription.Medicines = string(bytes)
		}
	}

	return dbPrescription
}