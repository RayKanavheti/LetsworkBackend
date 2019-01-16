package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/raykanavheti/LetsworkBackend/controllers/util"
	"github.com/raykanavheti/LetsworkBackend/models"
)

//EducationController interface
type EducationController struct{}

// CreateEducation creates a new education for a user
func (catCntrl *EducationController) CreateEducation(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()
	education := models.Education{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&education)
	if err != nil {
		mapError := map[string]string{"message": err.Error()}
		errj, _ := json.Marshal(mapError)
		responseWriter.WriteHeader(400)
		responseWriter.Write(errj)
	} else {
		valid := validation.Validation{}
		b, err := valid.Valid(education)
		if !b {
			mapError := map[string]string{"message": err.Error()}
			errj, _ := json.Marshal(mapError)
			responseWriter.WriteHeader(400)
			responseWriter.Write(errj)
		} else {
			cat, err := models.CreateEducation(education)
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
