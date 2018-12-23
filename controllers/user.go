package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gorilla/mux"
	"github.com/raykanavheti/LetsworkBackend/controllers/util"
	"github.com/raykanavheti/LetsworkBackend/models"
)

//UserController interface
type UserController struct{}

// CreateUser creates a new User resource
func (catCntrl *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		mapError := map[string]string{"message": err.Error()}
		errj, _ := json.Marshal(mapError)
		responseWriter.WriteHeader(400)
		responseWriter.Write(errj)
	} else {
		valid := validation.Validation{}
		b, err := valid.Valid(user)
		if !b {
			mapError := map[string]string{"message": err.Error()}
			errj, _ := json.Marshal(mapError)
			responseWriter.WriteHeader(400)
			responseWriter.Write(errj)
		} else {
			cat, err := models.CreateUser(user)
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

//GetAllUsers gets all users
func (catCntrl *UserController) GetAllUsers(w http.ResponseWriter, req *http.Request) {
	users, err := models.GetUsers()
	if err == nil {
		w.Header().Add("Content Type", "application/json")
		responseWriter := util.GetResponseWriter(w, req)
		defer responseWriter.Close()
		data, err := json.Marshal(users)
		if err == nil {
			responseWriter.Write(data)
		} else {
			errj, _ := json.Marshal(err)
			responseWriter.WriteHeader(404)
			responseWriter.Write(errj)
		}
	}
}

//GetUserByID for getting user by ID
func (catCntrl *UserController) GetUserByID(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	vars := mux.Vars(req)
	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw)
	if err == nil {
		user, err1 := models.GetUserByID(id)
		if err1 == nil {
			data, err2 := json.Marshal(user)
			if err2 == nil {
				responseWriter.Header().Add("Content Type", "application/json")
				responseWriter.Write(data)
			} else {
				responseWriter.WriteHeader(404)
				responseWriter.Write([]byte(err2.Error()))
			}
		} else {
			responseWriter.WriteHeader(404)
			responseWriter.Write([]byte(err1.Error()))
		}
	} else {
		responseWriter.WriteHeader(400)
		responseWriter.Write([]byte(err.Error()))
	}
}

// GetUserByUUID for getting user by ID
func (catCntrl *UserController) GetUserByUUID(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	vars := mux.Vars(req)
	uuidRaw := vars["uuid"]
	user, err := models.GetUserByUUID(uuidRaw)
	if err == nil {

		verifiedData, err2 := models.UpdateVerificationField(user)
		if err2 == nil {
			uj, _ := json.Marshal(verifiedData)
			responseWriter.Header().Add("Content Type", "application/json")
			responseWriter.Write(uj)
		} else {
			mapError := map[string]string{"message": err2.Error()}
			errj, _ := json.Marshal(mapError)
			responseWriter.WriteHeader(400)
			responseWriter.Write(errj)
		}

	} else {
		responseWriter.WriteHeader(404)
		responseWriter.Write([]byte(err.Error()))
	}

}

// SendRestLink It sends a link for a user to reset their password
func (catCntrl *UserController) SendRestLink(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	var buffer bytes.Buffer
	defer responseWriter.Close()
	vars := mux.Vars(req)
	userEmail := vars["email"]
	_, err := models.GetUserByEmail(userEmail)

	if err == nil {

		invokeSendLinkStatus := models.ResetPassword(userEmail)
		if invokeSendLinkStatus == "success" {
			responseWriter.Header().Add("Content Type", "application/json")
			buffer.WriteString(`{Response: "success", Message: "Email was sent successfully, Check your email !!!"}`)

			s := json.NewEncoder(w).Encode(buffer.String())
			r, _ := json.Marshal(s)
			responseWriter.Write(r)

		} else {
			responseWriter.Header().Add("Content Type", "application/json")
			buffer.WriteString(`{Response: "error", Message: "Sending email failed !!!"}`)

			s := json.NewEncoder(w).Encode(buffer.String())
			r, _ := json.Marshal(s)
			responseWriter.Write(r)
		}

	} else {
		responseWriter.WriteHeader(404)
		responseWriter.Write([]byte(err.Error()))
	}

}
