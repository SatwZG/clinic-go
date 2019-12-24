package api

import (
	"log"
	"math/rand"
	"net/http"
	"sync"

	"clinic-go/db"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var Token2ID sync.Map

func GetToken() string {
	b := make([]rune, 10)
	for {
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		_, ok := Token2ID.Load(string(b))
		if ok == false {
			break
		}
	}
	return string(b)
}

// status: (0, 1, 2) = (none, doctor, admin)
func DealwithCookie(r *http.Request) clinicDB.Account {
	cookie, err := r.Cookie("token")
	if err != nil {
		log.Println("读取cookie失败: ", err.Error())
	}
	ID, ok := Token2ID.Load(cookie.Value)
	if !ok {
		log.Println("cookie not exist: ", err.Error())
	}

	account := clinicDB.GetAccountByID(ID.(int))
	return account
}


