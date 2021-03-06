package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	uuid "github.com/satori/go.uuid"
	"github.com/urfave/negroni"
)

func newUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if r.Body == nil {
		http.Error(w, "no request body found", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	u := NewUser(
		user.ID,
		user.Name,
		user.Email,
	)
	saveErr := u.Save()
	if saveErr != nil {
		http.Error(w, "error while saving", 400)
		return
	}
}

func fetchOneUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	userID := vars["id"]
	user.ID = uuid.FromStringOrNil(userID)
	u := user.Fetch()
	json.NewEncoder(w).Encode(u)
}

func fetchAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	people := user.FetchAll()
	json.NewEncoder(w).Encode(people)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	id := vars["id"]
	user.ID = uuid.FromStringOrNil(id)
	err := user.Delete()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "user successfully deleted",
	})
}

func logger(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	t1 := time.Now()
	next(w, r)
	t2 := time.Now()
	log.Printf("[%s] %q %v", r.Method, r.URL.String(), t2.Sub(t1))
}

// Router is a collection of the application's routes and returns an instance of negroni
func Router() *negroni.Negroni {

	// router
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	n := negroni.New()
	n.Use(negroni.HandlerFunc(logger))
	n.Use(c)

	r.HandleFunc("/users", fetchAllUsersHandler).Methods("GET")
	r.HandleFunc("/users", newUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", fetchOneUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", deleteUserHandler).Methods("DELETE")

	n.UseHandler(r)
	return n
}

func main() {
	// init DB
	var s Storage
	s.BucketInit()

	fmt.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", Router()))
}
