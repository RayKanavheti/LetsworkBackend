package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/raykanavheti/LetsworkBackend/controllers/util"
	"github.com/raykanavheti/LetsworkBackend/models"
)

//AddressController interface
type AddressController struct{}

// CreateAddress creates a new address for a user
func (catCntrl *AddressController) CreateAddress(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()
	address := models.Address{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&address)
	if err != nil {
		mapError := map[string]string{"message": err.Error()}
		errj, _ := json.Marshal(mapError)
		responseWriter.WriteHeader(400)
		responseWriter.Write(errj)
	} else {
		valid := validation.Validation{}
		b, err := valid.Valid(address)
		if !b {
			mapError := map[string]string{"message": err.Error()}
			errj, _ := json.Marshal(mapError)
			responseWriter.WriteHeader(400)
			responseWriter.Write(errj)
		} else {
			cat, err := models.CreateAddress(address)
			if err == nil {
				uj, _ := json.Marshal(cat)
				responseWriter.Header().Set("Content-Type", "application/json")
				responseWriter.WriteHeader(201)
				responseWriter.Write(uj)
			} else {
				mapError := map[string]string{"message": err.Error()}
				errj, _ := json.Marshal(mapError)
				responseWriter.WriteHeader(400)
				responseWriter.Write(errj)
			}
		}
	}
}
