package middleware

import (
	"log"
	"net/http"
)

type AuthMiddleware func(w http.ResponseWriter, r *http.Request, uuid string)

func (m AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m(w, r, "")
}

// func AuthMiddleware(next AuthorizedHandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("auth middleware called for %s method and %s url", r.Method, r.URL.Path)
// 		next(w, r,"")
// 	})
// }

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request %s on path %s", r.Method, r.URL.Path)
		next(w, r)
	})
}
