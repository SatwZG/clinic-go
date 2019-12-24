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
	router.HandleFunc("/get_op_type", getOpType).Methods("POST")
	router.HandleFunc("/search_doctors", searchDoctors).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("start login")
	var a api.LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&a)
	if err != nil {
		log.Printf("", err)
		w.WriteHeader(500)
		return
	}
	log.Printf("ar: %s %s\n", a.Username, a.Password)
	item := clinicDB.GetAccountByUsername(a.Username)

	if item.Username != a.Username || item.Password != a.Password {
		w.WriteHeader(500)
		return
	}

	//cookie := http.Cookie{Name: "token", Value: "nill", Expires: expiration}
	cookie := http.Cookie {
		Name: "token",
		Value: api.GetToken(),
		MaxAge: 300,
	}
	api.Token2ID.Store(cookie.Value, item.ID)


	http.SetCookie(w, &cookie)
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
		w.WriteHeader(500)
		return
	}

	if len(req.Age) < 2 || req.Age[0] > req.Age[1] {
		log.Println("illegal query")
		w.WriteHeader(500)
		return
	}
	log.Printf("ar: %s %s %s %s %d\n", req.Name, req.Sex, req.Age, req.Department, req.Page)

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

	var doctors []api.Doctor

	for i := 0; i < len(items); i++ {
		doctors[i].ID = items[i].ID
		doctors[i].Username = items[i].Username
		doctors[i].Name = items[i].Name
		doctors[i].Department = items[i].Department
		doctors[i].Sex = items[i].Sex
		doctors[i].Age = items[i].Age
		doctors[i].Introduction = items[i].Introduction.String
		doctors[i].Avatar = items[i].Avatar.String
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