package main

import (
	"encoding/json"
	"net/http"

	"clinic-go/api"
	"clinic-go/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	clinicDB.InitDB()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/logout", logout).Methods("POST")
	router.HandleFunc("/get_op_type", getOpType).Methods("POST")

	router.HandleFunc("/search_doctors", searchDoctors).Methods("POST")
	router.HandleFunc("/add_doctor", addDoctor).Methods("POST")
	router.HandleFunc("/update_doctor", updateDoctor).Methods("POST")
	router.HandleFunc("/delete_doctor", deleteDoctor).Methods("POST")

	router.HandleFunc("/search_medicines_with_page", api.SearchMedicinesWithPage).Methods("POST")
	router.HandleFunc("/search_medicine", api.SearchMedicine).Methods("POST")
	router.HandleFunc("/add_medicine", api.AddMedicine).Methods("POST")
	router.HandleFunc("/update_medicine", api.UpdateMedicine).Methods("POST")
	router.HandleFunc("/delete_medicine", api.DeleteMedicine).Methods("POST")

	router.HandleFunc("/search_prescriptions_with_page", api.SearchPrescriptionsWithPage).Methods("POST")
	//router.HandleFunc("/search_prescriptions", api.SearchPrescriptions).Methods("POST")
	router.HandleFunc("/add_prescription", api.AddPrescription).Methods("POST")
	router.HandleFunc("/update_prescription", api.UpdatePrescription).Methods("POST")
	router.HandleFunc("/delete_prescription", api.DeletePrescription).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	//originsOk := handlers.AllowedOrigins([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT"})
	credentialsOk := handlers.AllowCredentials()

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk, credentialsOk)(router)))
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("start login")

	var a api.LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&a)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}
	log.Printf("ar: %s %s\n", a.Username, a.Password)
	item := clinicDB.GetAccountByUsername(a.Username)

	if item.Username != a.Username || item.Password != a.Password {
		w.WriteHeader(500)
		return
	}

	cookie := http.Cookie {
		Name: "token",
		Value: api.GetToken(),
		Path: "/",
		MaxAge: 3600,
	}
	api.Token2ID.Store(cookie.Value, item.ID)
	http.SetCookie(w, &cookie)

	resp := api.LoginResponse {
		Token: cookie.Value,
	}
	bytes, err := json.Marshal(resp)
	_, err = w.Write(bytes)
	if err != nil {
		log.Printf("this is maybe fail ?")
		w.WriteHeader(500)
	}
	log.Println("now response have token, after maybe no")
}

func logout(w http.ResponseWriter, r *http.Request) {
	log.Println("start logout")

	account := api.DealwithCookie(r)
	if account.Type == 0 {
		log.Warn("getOpType fail")
		w.WriteHeader(500)
		return
	}

	cookie, _ := r.Cookie("token")
	api.Token2ID.Delete(cookie.Value)

	w.WriteHeader(200)
}

func getOpType(w http.ResponseWriter, r *http.Request) {
	log.Println("start getOpType")

	account := api.DealwithCookie(r)
	if account.Type == 0 {
		log.Warn("getOpType fail")
		w.WriteHeader(500)
		return
	}

	res := api.GetOpTypeResponse {
		Type: account.Type,
	}

	if account.Type == 1 {
		doctor := clinicDB.GetDoctorByID(account.ID)
		res.Doctor = api.DbDoctor2req(doctor)

	} else if account.Type == 2 {
		admin := clinicDB.GetAdminByID(account.ID)
		res.Admin = api.DbAdmin2req(admin)
	}

	bytes, err := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bytes)

	if err != nil {
		log.Printf("this is maybe fail ?")
		w.WriteHeader(500)
	}
}

func searchDoctors(w http.ResponseWriter, r *http.Request) {
	log.Println("start searchDoctors")

	account := api.DealwithCookie(r)
	if account.Type != 2 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req api.DoctorSearchRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}

	if (len(req.Age) != 0 && len(req.Age) != 2) || (len(req.Age) == 2 && req.Age[0] > req.Age[1]) {
		log.Println("age illegal query")
		w.WriteHeader(500)
		return
	}
	log.Printf("ar: %s %s %s %s %d\n", req.Name, req.Sex, req.Age, req.Department, req.Page)

	if req.Page < 1 {
		req.Page = 1
	}

	var res api.DoctorSearchResponse
	items := api.GetDoctorsBySearch(req)
	if len(items) == 0 && req.Page > 1 {
		req.Page = 1
		items = api.GetDoctorsBySearch(req)
		if len(items) != 0 {
			log.Printf("now query page count = 0, so return 1 page content")
			res.NowPage = 1
		}
	} else {
		res.NowPage = req.Page
	}


	res.TotalPage = api.GetDoctorsTotalPageBySearch(req)

	doctors := make([]api.Doctor, len(items))

	for i := 0; i < len(items); i++ {
		doctors[i] = api.DbDoctor2req(items[i])
	}
	res.Doctors = doctors

	bytes, err := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bytes)

	if err != nil {
		log.Printf("this is maybe fail ?")
		w.WriteHeader(500)
	}
}

func addDoctor(w http.ResponseWriter, r *http.Request) {
	log.Println("start addDoctor")

	account := api.DealwithCookie(r)
	if account.Type != 2 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req api.AddDoctorRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}

	dbDoctor := api.Req2DbDoctor(req.Doctor)
	ID := clinicDB.AddDoctor(dbDoctor, req.Password)
	if ID == 0 {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func updateDoctor(w http.ResponseWriter, r *http.Request) {
	log.Println("start updateDoctor")

	account := api.DealwithCookie(r)
	if account.Type != 2 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req api.UpdateDoctorRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}

	dbDoctor := api.Req2DbDoctor(req.Doctor)
	ID := clinicDB.UpdateDoctor(dbDoctor, req.Password)
	if ID == 0 {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func deleteDoctor(w http.ResponseWriter, r *http.Request) {
	log.Println("start deleteDoctor")

	account := api.DealwithCookie(r)
	if account.Type != 2 {
		log.Warn("cookie miss match with operator")
		w.WriteHeader(500)
		return
	}

	var req api.DeleteDoctorRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Warn("", err)
		w.WriteHeader(500)
		return
	}

	// i don't how get delete status, but it can delete
	clinicDB.DeleteDoctor(req.Username)

	w.WriteHeader(200)
}