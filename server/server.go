package server

import (
	"crypto/sha1"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

type Token struct {
	Token string `json:"token"`
}

// Run server.
func Serve() {
	handle()

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Handle requests.
func handle() {
	fs := http.FileServer(http.Dir("web/dist"))
	http.Handle("/", fs)
	log.Println("Static content")

	http.HandleFunc("/auth", getAuthToken)
	http.HandleFunc("/ws", HandleConnections)

	go HandleMessages()
}

func getAuthToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bytes, _ := ioutil.ReadAll(r.Body)
	user := User{}
	err := json.Unmarshal(bytes, &user)
	if err != nil {
		log.Println("Can not read body while authentication")
		log.Println(err)
	}

	h := sha1.New()
	h.Write([]byte(user.Name))

	token := Token{
		Token: string(h.Sum(nil)),
	}

	responseBody, _ := json.Marshal(token)
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}
