package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"clinic-go/api"
	//"clinic-go/clinicDB"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//func init() {
//	clinicDB.InitDB()
//}

func main() {
	b := make([]byte, 32)
	rand.Read(b)
	fmt.Printf("%s\n", base64.URLEncoding.EncodeToString(b))
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Printf("hello\n")

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
		w.WriteHeader(500)
		return
	}
	log.Printf("%s %s", a.Username, a.Password)
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "token", Value: "nill", Expires: expiration}
	http.SetCookie(w, &cookie)
	w.WriteHeader(200)

}