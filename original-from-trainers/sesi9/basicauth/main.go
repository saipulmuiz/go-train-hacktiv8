package main

import (
	"basicauth/middleware"
	"basicauth/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/students", ActionStudent)
	port := ":4444"

	log.Println("server runnint at port", port)
	http.ListenAndServe(port, nil)
}

func ActionStudent(rw http.ResponseWriter, r *http.Request) {

	if !middleware.Auth(r) {
		writeJsonResponse(rw, http.StatusUnauthorized, map[string]string{
			"error": "you must be log in",
		})
		return
	}

	if !middleware.OnlyGet(r) {
		writeJsonResponse(rw, http.StatusUnauthorized, map[string]string{
			"error": "only method get allowed",
		})
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		if !middleware.AllowSuperman(r) {
			writeJsonResponse(rw, http.StatusUnauthorized, map[string]string{
				"error": "only superman get allowed",
			})
			return
		}
		id8, err := strconv.Atoi(id)
		if err != nil {
			writeJsonResponse(rw, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		user := models.GetUserById(int8(id8))
		if user == nil {
			writeJsonResponse(rw, http.StatusNotFound, map[string]string{
				"error": "user with id " + id + " not found",
			})
			return
		}
		writeJsonResponse(rw, http.StatusOK, user)
		return
	}
	writeJsonResponse(rw, http.StatusOK, models.GetUsers())
}

func writeJsonResponse(rw http.ResponseWriter, status int, payload interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(payload)
}
