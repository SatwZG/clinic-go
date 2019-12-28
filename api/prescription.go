package api

import (
	clinicDB "clinic-go/db"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SearchPrescriptionsWithPage(w http.ResponseWriter, r *http.Request) {
	log.Println("start SearchPrescriptionsWithPage")

	account := DealwithCookie(r)
	if account.Type != 1 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req SearchPrescriptionsWithPageRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}
	log.Println("SearchPrescriptionsWithPage req: ", req)

	// maybe need dealwith Time

	if req.Page < 1 {
		req.Page = 1
	}

	var res SearchPrescriptionsWithPageResponse
	filter := Req2DbPrescriptionFilter(req)
	items := clinicDB.GetPrescriptionsByFilter(filter)
	if len(items) == 0 && req.Page > 1 {
		req.Page = 1
		items = clinicDB.GetPrescriptionsByFilter(filter)
		if len(items) != 0 {
			log.Printf("now query page count = 0, so return 1 page content")
			res.NowPage = 1
		}
	} else {
		res.NowPage = req.Page
	}


	res.TotalPage = clinicDB.GetPrescriptionsTotalPageByFilter(filter)

	prescriptions := make([]Prescription, len(items))

	for i := 0; i < len(items); i++ {
		prescriptions[i] = DbPrescription2Req(items[i])
	}

	res.Prescriptions = prescriptions


	bytes, err := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bytes)

	if err != nil {
		log.Printf("this is maybe fail ?")
		w.WriteHeader(500)
	}
}

func AddPrescription(w http.ResponseWriter, r *http.Request) {
	log.Println("start AddPrescription")

	account := DealwithCookie(r)
	if account.Type != 1 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req AddPrescriptionRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}
	log.Println("AddPrescription req: ", req)

	prs := req.Prescription

	if len(prs.Medicines) == 0 || len(prs.Department) == 0 || len(prs.DoctorName) == 0 || len(prs.PatientName) == 0 {
		log.Warn("prescription haven't medicines")
		w.WriteHeader(500)
		return
	}


	dbPrescription := Req2DbPrescription(req.Prescription)
	doctor := clinicDB.GetDoctorByID(account.ID)
	if doctor.ID == 0 {
		log.Warn("operator not exit db")
	} else {
		dbPrescription.DoctorName = doctor.Name
	}

	dbMedicines := Req2DbMedicines(req.Prescription.Medicines)

	ID := clinicDB.AddPrescription(dbPrescription, dbMedicines)
	if ID == 0 {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func UpdatePrescription(w http.ResponseWriter, r *http.Request) {
	log.Println("start UpdatePrescription")

	account := DealwithCookie(r)
	if account.Type != 1 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req UpdatePrescriptionRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}
	log.Println("UpdatePrescription req: ", req)

	if req.Prescription.ID == 0 {
		log.Warn("UpdatePrescription reqID is empty")
		w.WriteHeader(500)
		return
	}

	dbPrescription := Req2DbPrescription(req.Prescription)
	dbPrescription.ID = req.Prescription.ID
	ID := clinicDB.UpdatePrescription(dbPrescription)
	if ID == 0 {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func DeletePrescription(w http.ResponseWriter, r *http.Request) {
	log.Println("start DeletePrescription")

	account := DealwithCookie(r)
	if account.Type != 1 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req DeletePrescriptionRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}
	log.Println("DeletePrescription req: ", req)

	clinicDB.DeletePrescription(req.ID)

	w.WriteHeader(200)
}
