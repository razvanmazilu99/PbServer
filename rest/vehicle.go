package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"problem/entity"
)

func GetMyVehicle(rw http.ResponseWriter, r *http.Request) {

	vehicleType := r.URL.Query().Get("type")

	var vehicle entity.Vehicle

	switch vehicleType {
	case "car":
		vehicle = entity.Car{Name: "BMW", Type: "Car", Model: 2020}
	case "bike":
		vehicle = entity.Bike{Name: "BMX", Type: "Bike", Model: 2009}
	case "bus":
		vehicle = entity.Bus{Name: "STPT", Type: "Bus", Model: 2000}
	default:
		rw.Write([]byte("Vehicle not known"))
		return
	}

	vehicleBytes, err := json.Marshal(vehicle)

	if err == nil {
		rw.Write([]byte(string(vehicleBytes)))
	}

}

func PostMyVehicle(rw http.ResponseWriter, r *http.Request) {

	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)

	if err != nil {
		rw.Write([]byte("Error 1"))
		return
	}

	var vehicle entity.Vehicle
	//err = json.Unmarshal(bodyBytes, &vehicle)

	/*if err != nil {
		rw.Write([]byte("Error 2"))
		return
	}*/

	fmt.Println(vehicle)
	rw.Write(bodyBytes)
}
