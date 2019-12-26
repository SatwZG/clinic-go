package clinicDB

func GetMedicinesByName(name string) []Medicine {
	var medicines []Medicine
	DB.Table("medicines").
		Where("name like ?", "%"+name+"%").
		Find(&medicines)
	return medicines
}

func GetMedicineByName(name string) Medicine {
	var medicine Medicine
	DB.Table("medicines").
		Where("name like ?", name).
		Find(&medicine)
	return medicine
}

func AddMedicine(medicine Medicine) int {
	DB.Table("medicines").
		Create(&medicine)
	return medicine.ID
}

func UpdateMedicine(medicine Medicine) int {
	var dbMedicine Medicine
	DB.Table("medicines").
		Where("name = ?", medicine.Name).
		First(&dbMedicine)

	if dbMedicine.ID != 0 {
		dbMedicine.Count = medicine.Count
		DB.Table("medicines").
			Save(&dbMedicine)
	}
	return medicine.ID
}

func DeleteMedine(name string) {
	DB.Where("name = ?", name).
		Delete(Medicine{})
}
