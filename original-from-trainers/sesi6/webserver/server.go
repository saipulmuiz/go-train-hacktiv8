package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Id       string
	Name     string
	Age      int
	Division string
}

var users = []User{
	{
		Id:       "user-1",
		Name:     "User 1",
		Age:      20,
		Division: "IT",
	},
	{
		Id:       "user-2",
		Name:     "User 2",
		Age:      20,
		Division: "IT",
	},
	{
		Id:       "user-3",
		Name:     "User 3",
		Age:      20,
		Division: "IT",
	},
	{
		Id:       "user-4",
		Name:     "User 4",
		Age:      20,
		Division: "IT",
	},
}

type HTML struct {
	Title string
	Data  []User
}

func main() {

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/add", createUser)

	port := ":8001"

	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}

func getUsers(rw http.ResponseWriter, r *http.Request) {
	traceId := "91aaf539-3b07-42fd-bd05-713e7f5f7c06"
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("x-trace-id", traceId)
	header := r.Header.Get("client-id")

	if r.Method == http.MethodGet {
		log.Println("[INFO]", r.Method, r.URL.Path, "trace id :", traceId)
		if header == "API" {

			json.NewEncoder(rw).Encode(users)
			return
		} else {
			rw.Header().Set("Content-Type", "text/html")
			tpl, err := template.ParseFiles("./index.html")
			if err != nil {
				log.Println("[ERROR]", r.Method, r.URL.Path, "trace id :", traceId, "error :", err.Error())
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			d := HTML{
				Title: "halaman users",
				Data:  users,
			}

			tpl.Execute(rw, d)
			return
		}
	}

	errMsg := fmt.Sprintf("method %s not allowed !", r.Method)
	log.Println("[ERROR]", r.Method, r.URL.Path, "trace id :", traceId, "error :", errMsg)
	rw.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(rw).Encode(map[string]string{
		"error": "Ga boleh oi",
	})
}

func createUser(rw http.ResponseWriter, r *http.Request) {
	traceId := "91aaf539-3b07-42fd-bd05-713e7f5f7c06"
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("x-trace-id", traceId)

	if r.Method == http.MethodPost {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{
				"error": "bad request",
			})
			log.Println("[ERROR]", r.Method, r.URL.Path, "trace id :", traceId, "error :", err.Error())
			return
		}

		user.Id = fmt.Sprintf("user-%d", len(users)+1)

		users = append(users, user)
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(map[string]string{
			"message": "created success",
		})
		return
	}

	errMsg := fmt.Sprintf("method %s not allowed !", r.Method)
	log.Println("[ERROR]", r.Method, r.URL.Path, "trace id :", traceId, "error :", errMsg)
	rw.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(rw).Encode(map[string]string{
		"error": "Ga boleh oi",
	})
}
