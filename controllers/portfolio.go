package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/raykanavheti/LetsworkBackend/controllers/util"
	"github.com/raykanavheti/LetsworkBackend/models"
)

//PortfolioController interface
type PortfolioController struct{}

// Createportfolios creates a new portfolio for a user
func (catCntrl *PortfolioController) CreatePortfolios(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()
	portfolios := []models.Portfolio{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&portfolios)
	if err != nil {
		mapError := map[string]string{"message": err.Error()}
		errj, _ := json.Marshal(mapError)
		responseWriter.WriteHeader(400)
		responseWriter.Write(errj)
	} else {
			cat, err := models.CreatePortfolios(portfolios)
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
