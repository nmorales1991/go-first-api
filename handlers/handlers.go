package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nmorales1991/go-first-api/middlew"
	"github.com/nmorales1991/go-first-api/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Manejadores() {
	router := mux.NewRouter()
	//llamo a un post, pasa por un middleware para checkear el status de la BD, si pasa bien entonces llamo a routers RegistroUsuario
	router.HandleFunc("/registro", middlew.CheckBD(routers.RegistroUsuario)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	handler := cors.AllowAll().Handler(router)
	fmt.Println("Listening port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
