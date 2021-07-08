package rest

import (
	"encoding/json"
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
		rw.Write([]byte("Internal issue 1"))
		return
	}

	var vehicle map[string]interface{}
	err = json.Unmarshal(bodyBytes, &vehicle)

	if err != nil {
		rw.Write([]byte("Internal issue 2"))
		return
	}

	vehicleType, ok := vehicle["ype"].(string) //schimbam aici tipul in string ca sa nu punem .(type) la switch

	if !ok {
		rw.Write([]byte("Internal issue 3"))
		return
	}

	switch vehicleType {
	case "car":
		rw.Write([]byte("This is a car"))
	case "bike":
		rw.Write([]byte("This is a bike"))
	case "bus":
		rw.Write([]byte("This is a bus"))
	default:
		rw.Write([]byte("Vehicle not known"))
		return
	}

	rw.Write(bodyBytes)
}
