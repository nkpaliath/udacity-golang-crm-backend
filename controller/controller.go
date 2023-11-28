package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nkpalaith/udacity-golang-crm-backend/db"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.Customers)
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	var newCustomer db.Customer
	newId := len(db.Customers) + 1
	data, _ := io.ReadAll(r.Body)
	json.Unmarshal(data, &newCustomer)
	newCustomer.ID = newId
	db.Customers[newId] = newCustomer

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(db.Customers[newId])
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	w.Header().Set("Content-Type", "application/json")
	if c, ok := db.Customers[id]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(c)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[int]db.Customer{})
	}
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var updatedCustomer db.Customer
	updatedCustomer.ID = id
	data, _ := io.ReadAll(r.Body)
	json.Unmarshal(data, &updatedCustomer)

	w.Header().Set("Content-Type", "application/json")
	if customer, ok := db.Customers[id]; ok {

		if updatedCustomer.Contacted != nil {
			customer.Contacted = updatedCustomer.Contacted
		}

		if updatedCustomer.Email != "" {

			customer.Email = updatedCustomer.Email
		}
		if updatedCustomer.Name != "" {
			customer.Name = updatedCustomer.Name
		}
		if updatedCustomer.Role != "" {
			customer.Role = updatedCustomer.Role
		}
		if updatedCustomer.Phone != "" {
			customer.Phone = updatedCustomer.Phone
		}
		db.Customers[id] = customer
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(db.Customers[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[int]db.Customer{})
	}
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	w.Header().Set("Content-Type", "application/json")
	if c, ok := db.Customers[id]; ok {
		delete(db.Customers, c.ID)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[int]db.Customer{})
	}
}
