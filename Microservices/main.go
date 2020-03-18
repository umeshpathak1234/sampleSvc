package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//Account is used to hold the account details.
type Account struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber int64  `json:"phonenumber"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Status      string `json:"status"`
}

//acc will hold all the account details
var acc = make(map[string]Account)

func main() {
	intiHandler()

}
func intiHandler() {
	r := mux.NewRouter()
	r.HandleFunc("/account", createAccountHandler).Methods(http.MethodPost)
	r.HandleFunc("/account", getAccountHandler).Methods(http.MethodGet)
	r.HandleFunc("/account", updateAccountHandler).Methods(http.MethodPut)
	r.HandleFunc("/account", deleteAccountHandler).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", r)

}
func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	var a Account

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	// add new id to the new account

	a.ID = uuid.New().String()
	a.Status = "ACTIVE"
	// store account details in map
	//map key is ID, value is Account
	acc[a.ID] = a

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Account created succefully")
}
func getAccountHandler(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.URL.Query()["id"]; ok {
		if err := json.NewEncoder(w).Encode(acc[id[0]]); err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		if err := json.NewEncoder(w).Encode(&acc); err != nil {
			w.Write([]byte(err.Error()))
		}
	}

	fmt.Fprintln(w, "Ahile lai checking matrai ho hai")
}
func updateAccountHandler(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.URL.Query()["id"]; ok {
		var a Account
		// decoding the payload of the request
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		//update the id of the account
		a.ID = id[0]
		//apply the update
		acc[id[0]] = a
		a.Status = "ACTIVE"

		w.WriteHeader(http.StatusNoContent)
		return

	}
	w.WriteHeader(http.StatusBadRequest)
	return
}
func deleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.URL.Query()["id"]; ok {
		//delete the account matching with this id
		a := acc[id[0]]
		a.Status = "DELETED"

		// store account back to the map
		acc[id[0]] = a

		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}
