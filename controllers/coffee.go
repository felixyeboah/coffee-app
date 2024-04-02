package controllers

import (
	"coffee-app/helpers"
	"coffee-app/services"
	"encoding/json"
	"github.com/go-chi/chi/v5"
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

func GetCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	coffee, err := coffeeService.GetByID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.Envelop{"error": "Internal Server Error"})
		return
	}

	helpers.WriteJson(w, http.StatusOK, helpers.Envelop{"data": coffee})
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

func UpdateCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	coffee := services.Coffee{}
	err := json.NewDecoder(r.Body).Decode(&coffee)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJson(w, http.StatusBadRequest, helpers.Envelop{"error": "Invalid request payload"})
		return
	}

	updatedCoffee, err := coffeeService.Update(id, coffee)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.Envelop{"error": "Internal Server Error"})
		return
	}

	helpers.WriteJson(w, http.StatusOK, helpers.Envelop{"data": updatedCoffee})
}

func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := coffeeService.Delete(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJson(w, http.StatusInternalServerError, helpers.Envelop{"error": "Internal Server Error"})
		return
	}

	helpers.WriteJson(w, http.StatusNoContent, nil)
}
