package controllers

import (
	"coffee-app/helpers"
	"coffee-app/services"
	"encoding/json"
	"net/http"
)

var coffeeService services.Coffee

// GetCoffees GET /coffees
func GetCoffees(w http.ResponseWriter, r *http.Request) {
	coffees, err := coffeeService.GetAll()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.Envelop{"error": "Internal Server Error"})
		return
	}

	helpers.WriteJson(w, http.StatusOK, helpers.Envelop{"data": coffees})
}

// CreateCoffee POST /coffees
func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	coffee := services.Coffee{}
	err := json.NewDecoder(r.Body).Decode(&coffee)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJson(w, http.StatusBadRequest, helpers.Envelop{"error": "Invalid request payload"})
		return
	}

	createdCoffee, err := coffeeService.Create(coffee)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.Envelop{"error": "Internal Server Error"})
		return
	}

	helpers.WriteJson(w, http.StatusCreated, helpers.Envelop{"data": createdCoffee})
}
