package middlew

import (
	"github.com/nmorales1991/go-first-api/bd"
	"net/http"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if bd.Ping() == false {
			http.Error(writer, "Conexión perdida", 500)
		}
		next.ServeHTTP(writer, request)
	}
}
