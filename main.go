package main

import (
	"net/http"
	"problem/rest"

	"github.com/go-chi/chi"
)

func main() {

	http.HandleFunc("/", rest.GetMyVehicle)

	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Get("/getmyvehicle", rest.GetMyVehicle)
		r.Post("/postmyvehicle", rest.PostMyVehicle)
		r.Post("/postcandrive", rest.PostCanDrive)
	})

	http.ListenAndServe(":8080", router)
}
