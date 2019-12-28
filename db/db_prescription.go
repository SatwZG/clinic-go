package clinicDB

import (
	"log"
	"time"
)

type PrescriptionFilter struct {
	Department string
	DoctorName string
	PatientName string
	Age []int
	Sex string
	Time []time.Time
	Page int
}

func GetPrescriptionsByFilter(filter PrescriptionFilter) []Prescription {
	var prescriptions []Prescription

	if len(filter.Age) != 2 {
		filter.Age = []int{0, 1000000}
	} else if filter.Age[0] > filter.Age[1] {
		t := filter.Age[0]
		filter.Age[0] = filter.Age[1]
		filter.Age[1] = t
	}

	if len(filter.Time) != 2 {
		filter.Time = make([]time.Time, 2)
		filter.Time[0] = time.Unix(1469579899, 0)
		filter.Time[1] = time.Now()
	} else if filter.Time[1].Sub(filter.Time[0]).Nanoseconds() < 0 {
		t := filter.Time[0]
		filter.Time[0] = filter.Time[1]
		filter.Time[1] = t
	}

	log.Println("#", filter.Time[0])
	log.Println("$", filter.Time[1])


	DB.Table("prescriptions").
		Where("created_at BETWEEN ? AND ?", filter.Time[0], filter.Time[1]).
		Where("doctor_name like ? AND patient_name like ?", "%"+filter.DoctorName+"%", "%"+filter.PatientName+"%").
		Where("department like ? AND sex like ?", "%"+filter.Department+"%", "%"+filter.Sex+"%").
		Where("age >= ? AND age <= ?", filter.Age[0], filter.Age[1]).
		Limit(PageMax).
		Offset((filter.Page-1)*PageMax).
		Find(&prescriptions)

	return prescriptions
}

func GetPrescriptionsCountByFilter(filter PrescriptionFilter) int {
	var count int
	if len(filter.Age) == 0 {
		filter.Age = []int{0, 1000000}
	}
	DB.Table("prescriptions").
		Where("doctor_name like ? AND patient_name like ?", "%"+filter.DoctorName+"%", "%"+filter.PatientName+"%").
		Where("department like ? AND sex like ?", "%"+filter.Department+"%", "%"+filter.Sex+"%").
		Where("age >= ? AND age <= ?", filter.Age[0], filter.Age[1]).
		Count(&count)
	return count
}

func GetPrescriptionsTotalPageByFilter(filter PrescriptionFilter) int {
	count := GetPrescriptionsCountByFilter(filter)
	if count == 0 {
		return 1
	}

	if count%PageMax == 0 {
		return count/PageMax
	} else {
		return count/PageMax + 1
	}
}

func DeletePrescription(ID int) {
	DB.Where("id = ?", ID).
		Delete(Prescription{})
}

func AddPrescription(prescription Prescription, medicines []Medicine) int {
	tx := DB.Begin()

	dbMedicines := make([]Medicine, len(medicines))
	for i := 0; i < len(medicines); i++ {
		tx.Table("medicines").
			Where("name like ?", medicines[i].Name).
			Find(&dbMedicines[i])

		if dbMedicines[i].Count < medicines[i].Count {
			tx.Commit()
			return 0
		}
		dbMedicines[i].Count -= medicines[i].Count
	}
	log.Println("", dbMedicines)
	for i := 0; i < len(dbMedicines); i++ {
		tx.Table("medicines").
			Save(&dbMedicines[i])
	}

	tx.Table("prescriptions").
		Create(&prescription)

	tx.Commit()
	return prescription.ID
}

func UpdatePrescription(prescription Prescription) int {
	var dbPrescription Prescription
	DB.Table("prescriptions").
		Where("id = ?", prescription.ID).
		First(&dbPrescription)
	if dbPrescription.ID != 0 {
		if prescription.PatientName != "" {
			dbPrescription.PatientName = prescription.PatientName
		}
		if prescription.Age != 0 {
			dbPrescription.Age = prescription.Age
		}
		if prescription.Sex != "" {
			dbPrescription.Sex = prescription.Sex
		}
		DB.Table("prescriptions").
			Save(&dbPrescription)
	}
	return dbPrescription.ID
}