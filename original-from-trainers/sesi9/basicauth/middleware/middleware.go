package middleware

import (
	"net/http"
)

var (
	ACCOUNT map[string]string = map[string]string{
		"batman":   "superman",
		"superman": "manusia-super",
	}
)

var ALLOWED_ROLE map[string]string = map[string]string{
	"batman":   "get all",
	"superman": "*",
}

func Auth(r *http.Request) (isAllow bool) {
	username, password, ok := r.BasicAuth()
	if !ok {
		return
	}

	return ACCOUNT[username] == password
}

func AllowBatman(r *http.Request) bool {
	username, _, _ := r.BasicAuth()
	return allow(username, "get all")
}

func AllowSuperman(r *http.Request) bool {
	username, _, _ := r.BasicAuth()
	return allow(username, "*")
}

func allow(role, access string) bool {
	val, ok := ALLOWED_ROLE[role]
	if !ok {
		return false
	}

	return val == access
}

func OnlyGet(r *http.Request) bool {
	return onlyMethod(r.Method, "GET")
}

func onlyMethod(method, rule string) bool {
	return method == rule
}
