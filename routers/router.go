package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raykanavheti/LetsworkBackend/controllers"
)

var (
	user      = new(controllers.UserController)
	profile   = new(controllers.ProfileController)
	education = new(controllers.EducationController)
	portfolio = new(controllers.PortfolioController)
	address   = new(controllers.AddressController)
	project   = new(controllers.ProjectController)
	skill     = new(controllers.SkillController)
)

//InitRoutes : Registering all system Routes.
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(201)
	})
	router.HandleFunc("/api/user", user.CreateUser).Methods("POST")
	router.HandleFunc("/api/user_update", user.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users", user.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", user.GetUserByID).Methods("GET")
	router.HandleFunc("/api/user_confirm/{uuid}", user.GetUserByUUID).Methods("GET")
	router.HandleFunc("/api/reset_password/{email}", user.SendRestLink).Methods("GET")

	router.HandleFunc("/api/profile", profile.CreateProfile).Methods("POST")

	router.HandleFunc("/api/educations", education.CreateEducations).Methods("POST")

	router.HandleFunc("/api/portfolios", portfolio.CreatePortfolios).Methods("POST")

	router.HandleFunc("/api/address", address.CreateAddress).Methods("POST")

	router.HandleFunc("/api/project", project.CreateProject).Methods("POST")
	router.HandleFunc("/api/project_update", project.UpdateProject).Methods("PUT")

	router.HandleFunc("/api/skills", skill.CreateSkills).Methods("POST")

	return router
}
