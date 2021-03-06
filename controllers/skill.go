package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/raykanavheti/LetsworkBackend/controllers/util"
	"github.com/raykanavheti/LetsworkBackend/models"
)

//SkillController interface
type SkillController struct{}

// CreateSkills creates a new Skill for a skill
func (catCntrl *SkillController) CreateSkills(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()
	Skills := []models.Skill{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&Skills)
	if err != nil {
		mapError := map[string]string{"message": err.Error()}
		errj, _ := json.Marshal(mapError)
		responseWriter.WriteHeader(400)
		responseWriter.Write(errj)
	} else {
			cat, err := models.CreateSkills(Skills)
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

//GetAllSkills gets all Skills
func (catCntrl *SkillController) GetAllSkills(w http.ResponseWriter, req *http.Request) {
	skills, err := models.GetSkills()
	if err == nil {
		w.Header().Add("Content Type", "application/json")
		responseWriter := util.GetResponseWriter(w, req)
		defer responseWriter.Close()
		data, err := json.Marshal(skills)
		if err == nil {
			responseWriter.Write(data)
		} else {
			errj, _ := json.Marshal(err)
			responseWriter.WriteHeader(404)
			responseWriter.Write(errj)
		}
	}
}
