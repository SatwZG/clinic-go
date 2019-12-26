package api

import (
	"encoding/json"
	"net/http"

	clinicDB "clinic-go/db"
	log "github.com/sirupsen/logrus"
)

func SearchMedicinesWithPage(w http.ResponseWriter, r *http.Request) {
	log.Println("start SearchMedicinesWithPage")

	account := DealwithCookie(r)
	if account.Type != 2 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req SearchMedicineWithPageRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}
	log.Println("SearchMedicinesWithPage req: ", req)

	if req.Page < 1 {
		req.Page = 1
	}

	var resp SearchMedicineWithPageResponse
	items := clinicDB.GetMedicinesByNamePage(req.Name, req.Page)
	if len(items) == 0 && req.Page > 1 {
		req.Page = 1
		items = clinicDB.GetMedicinesByNamePage(req.Name, req.Page)
		if len(items) != 0 {
			log.Printf("now query page count = 0, so return 1 page content")
			resp.NowPage = 1
		}
	} else {
		resp.NowPage = req.Page
	}


	resp.TotalPage = clinicDB.GetMedicinesTotalPageByName(req.Name)

	medicines := make([]Medicine, len(items))
	for i := 0; i < len(items); i++ {
		medicines[i] = DbMedicine2Req(items[i])
	}
	resp.Medicines = medicines

	bytes, err := json.Marshal(resp)
	_, err = w.Write(bytes)

	if err != nil {
		log.Printf("this is maybe fail ?")
		w.WriteHeader(500)
	}
}

func SearchMedicine(w http.ResponseWriter, r *http.Request) {
	log.Println("start SearchMedicine")

	account := DealwithCookie(r)
	if account.Type != 2 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req SearchMedicineRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}
	log.Println("SearchMedicine req: ", req)

	var resp SearchMedicineResponse
	items := clinicDB.GetMedicinesByName(req.Name)

	medicines := make([]Medicine, len(items))
	for i := 0; i < len(items); i++ {
		medicines[i] = DbMedicine2Req(items[i])
	}
	resp.Medicines = medicines

	bytes, err := json.Marshal(resp)
	_, err = w.Write(bytes)

	if err != nil {
		log.Printf("this is maybe fail ?")
		w.WriteHeader(500)
	}
}

func AddMedicine(w http.ResponseWriter, r *http.Request) {
	log.Println("start AddMedicine")

	account := DealwithCookie(r)
	if account.Type != 2 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req AddMedicineRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}
	log.Println("AddMedicine req: ", req)

	dbMedicine := Req2DbMedicine(req.Medicine)
	ID := clinicDB.AddMedicine(dbMedicine)
	if ID == 0 {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func UpdateMedicine(w http.ResponseWriter, r *http.Request) {
	log.Println("start UpdateMedicine")

	account := DealwithCookie(r)
	if account.Type != 2 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req UpdateMedicineRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}
	log.Println("UpdateMedicine req: ", req)

	dbMedicine := Req2DbMedicine(req.Medicine)
	ID := clinicDB.UpdateMedicine(dbMedicine)
	if ID == 0 {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func DeleteMedicine(w http.ResponseWriter, r *http.Request) {
	log.Println("start DeleteMedicine")

	account := DealwithCookie(r)
	if account.Type != 2 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req DeleteMedicineRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}
	log.Println("DeleteMedicine req: ", req)

	clinicDB.DeleteMedine(req.Name)

	w.WriteHeader(200)
}
