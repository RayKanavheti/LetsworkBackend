package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gorilla/mux"
	"github.com/raykanavheti/LetsworkBackend/controllers/util"
	"github.com/raykanavheti/LetsworkBackend/models"
)

// ProjectController interface
type ProjectController struct{}

// CreateProject creates a new project
func (ProjCtrl *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)

	//	file, handle, err := r.FormFile("file")
	defer responseWriter.Close()
	project := models.Project{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&project)
	fmt.Printf("Project %v", r.Body)
	if err != nil {
		mapError := map[string]string{"message": err.Error()}
		errj, _ := json.Marshal(mapError)
		responseWriter.WriteHeader(400)
		responseWriter.Write(errj)
	} else {
		valid := validation.Validation{}
		b, err := valid.Valid(project)
		if !b {
			mapError := map[string]string{"message": err.Error()}
			errj, _ := json.Marshal(mapError)
			responseWriter.WriteHeader(400)
			responseWriter.Write(errj)
		} else {
			cat, err := models.CreateProject(project)
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

// UpdateProject updates a project
func (ProjCtrl *ProjectController) UpdateProject(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()
	project := models.Project{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&project)

	if err != nil {
		mapError := map[string]string{"message": err.Error()}
		errj, _ := json.Marshal(mapError)
		responseWriter.WriteHeader(400)
		responseWriter.Write(errj)
	} else {
		valid := validation.Validation{}
		b, err := valid.Valid(project)
		if !b {
			mapError := map[string]string{"message": err.Error()}
			errj, _ := json.Marshal(mapError)
			responseWriter.WriteHeader(400)
			responseWriter.Write(errj)
		} else {
			cat, err := models.UpdateProject(project)
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

//GetProjectsByStatusAndOwnerId gets all projects related to the owner depending on the status of the project
func (ProjCtrl *ProjectController) GetProjectsByStatusAndOwnerId(w http.ResponseWriter, req *http.Request) {
		responseWriter := util.GetResponseWriter(w, req)
	vars := mux.Vars(req)
	OwnerIDRaw := vars["OwnerID"]
	Status := vars["Status"]
	OwnerID, err1 := strconv.Atoi(OwnerIDRaw)
	if err1 == nil {
	projects, err := models.GetProjectsByStatusAndOwnerID(OwnerID, Status)
	if err == nil {
		w.Header().Add("Content Type", "application/json")
		defer responseWriter.Close()
		data, err := json.Marshal(projects)
		if err == nil {
			responseWriter.Write(data)
		} else {
			errj, _ := json.Marshal(err)
			responseWriter.WriteHeader(404)
			responseWriter.Write(errj)
		}
	}
} else {
	responseWriter.WriteHeader(400)
	responseWriter.Write([]byte(err1.Error()))
}
}

//GetallProjectsByStatus gets all projects depending to the status of the project
func (ProjCtrl *ProjectController) GetallProjectsByStatus(w http.ResponseWriter, req *http.Request) {
		responseWriter := util.GetResponseWriter(w, req)
	vars := mux.Vars(req)
	Status := vars["Status"]
	if err1 == nil {
	projects, err := models.GetAllProjectsByStatus(Status)
	if err == nil {
		w.Header().Add("Content Type", "application/json")
		defer responseWriter.Close()
		data, err := json.Marshal(projects)
		if err == nil {
			responseWriter.Write(data)
		} else {
			errj, _ := json.Marshal(err)
			responseWriter.WriteHeader(404)
			responseWriter.Write(errj)
		}
	}
} else {
	responseWriter.WriteHeader(400)
	responseWriter.Write([]byte(err1.Error()))
}
}
