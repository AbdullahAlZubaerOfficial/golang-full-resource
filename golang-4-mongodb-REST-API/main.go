package main

import (
	"net/http"
	"time"

	"github.com/abdullahalzubaerofficial/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	session := getSession()
	defer session.Close()

	uc := controllers.NewUserController(session)

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe(":9000", r)
}

func getSession() *mgo.Session {
    s, err := mgo.DialWithTimeout("mongodb+srv://username:password@cluster0.rapspdi.mongodb.net/?appName=Cluster0", 10*time.Second)
    if err != nil {
        panic(err)  
    }
    s.SetMode(mgo.Monotonic, true)
    return s
}

