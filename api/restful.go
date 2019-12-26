package api

import "time"

type LoginRequest struct {
	Username string  `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Admin struct {
	Username string 	`json:"username"`
	Name string 		`json:"name"`
}

type Doctor struct {
	Username string 	`json:"username"`
	Name string 		`json:"name"`
	Sex string 			`json:"sex"`
	Age int 			`json:"age"`
	Department string   `json:"department"`
	Avatar string       `json:"avatar"`
	Introduction string `json:"introduction"`
}

type GetOpTypeRequest struct {
}
// status: (0, 1, 2) = (none, doctor, admin)
type GetOpTypeResponse struct {
	Doctor Doctor `json:"doctor"`
	Admin Admin   `json:"admin"`
	Type int	  `json:"type"`
}


type DoctorSearchRequest struct {
	Name string 	  `json:"name"`
	Department string `json:"department"`
	Sex string 		  `json:"sex"`
	Age []int 		  `json:"age"`
	Page int 		  `json:"page"`
}

type DoctorSearchResponse struct {
	NowPage int	 	 `json:"nowPage"`
	TotalPage int	 `json:"totalPage"`
	Doctors []Doctor `json:"doctors"`
}

type AddDoctorRequest struct {
	Doctor Doctor 	`json:"doctor"`
	Password string `json:"password"`
}
type AddDoctorResponse struct {
}

type UpdateDoctorRequest struct {
	Doctor Doctor 	`json:"doctor"`
	Password string `json:"password"`
}
type UpdateDoctorResponse struct {

}

type DeleteDoctorRequest struct {
	Username string `json:"username"`
}
type DeleteDoctorResponse struct {

}

// Medicine
type Medicine struct {
	Name  string `json:"name"`
	Count int `json:"count"`
}
type SearchMedicineRequest struct {
	Name  string `json:"name"`

}
type SearchMedicineResponse struct {
	Medicines []Medicine `json:"medicines"`
}

type AddMedicineRequest struct {
	Medicine Medicine `json:"medicine"`
}
type AddMedicineResponse struct {

}

type DeleteMedicineRequest struct {
	Name  string `json:"name"`
}
type DeleteMedicineResponse struct {

}

type UpdateMedicineRequest struct {
	Medicine Medicine `json:"medicine"`
}
type UpdateMedicineResponse struct {

}

// Prescription
type Prescription struct {
	ID int               `json:"id"`
	Department string    `json:"department"`
	DoctorName string	 `json:"doctorName"`
	PatientName string   `json:"patientName"`
	Age int              `json:"age"`
	Sex string           `json:"sex"`
	Medicines []Medicine `json:"medicines"`
	CreateTime time.Time `json:"createTime"`
}

type SearchPrescriptionsWithPageRequest struct {
	Department string   `json:"department"`
	DoctorName string	`json:"doctorName"`
	PatientName string  `json:"patientName"`
	Age []int           `json:"age"`
	Sex string          `json:"sex"`
	Time []time.Time 	`json:"time"`
	Page int            `json:"page"`
}

type SearchPrescriptionsWithPageResponse struct {
	NowPage int	 	 			 `json:"nowPage"`
	TotalPage int	 			 `json:"totalPage"`
	Prescriptions []Prescription `json:"prescriptions"`
	Test time.Time               `json:"test"`
}

//type SearchPrescriptionsRequest struct {
//	Name  string `json:"name"`
//
//}
//type SearchPrescriptionsResponse struct {
//	Medicines []Medicine `json:"medicines"`
//}

type AddPrescriptionRequest struct {
	Prescription Prescription `json:"prescription"`
}
type AddPrescriptionResponse struct {

}

type DeletePrescriptionRequest struct {
	ID int `json:"id"`
}
type DeletePrescriptionResponse struct {

}

type UpdatePrescriptionRequest struct {
	Prescription Prescription `json:"prescription"`
}
type UpdatePrescriptionResponse struct {

}

