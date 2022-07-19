package middleware

import "net/http"

const (
	USERNAME string = "batman"
	PASSWORD string = "superman"
)

func Auth(r *http.Request) (isAllow bool) {
	username, password, ok := r.BasicAuth()
	if !ok {
		return
	}

	return USERNAME == username && PASSWORD == password
}

func OnlyGet(r *http.Request) bool {
	return OnlyMethod(r.Method, "GET")
}

func OnlyMethod(method, rule string) bool {
	return method == rule
}

func OnlyUser(username string) bool {
	return username == USERNAME
}
