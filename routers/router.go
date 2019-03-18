package routers

import (
	// "net/http"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/raykanavheti/LetsworkBackend/controllers"
)

var (
	user        = new(controllers.UserController)
	profile     = new(controllers.ProfileController)
	education   = new(controllers.EducationController)
	portfolio   = new(controllers.PortfolioController)
	address     = new(controllers.AddressController)
	project     = new(controllers.ProjectController)
	skill       = new(controllers.SkillController)
	projectFile = new(controllers.ProjectFileController)
)

//InitRoutes : Registering all system Routes.
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, enctype, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.WriteHeader(201)
		}

	})
	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
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
	router.HandleFunc("/api/projects_status/{OwnerID}/{Status}", project.GetProjectsByStatusAndOwnerId).Methods("GET")

	router.HandleFunc("/api/files/{id}", projectFile.CreateProjectFiles).Methods("POST")

	router.HandleFunc("/api/skills", skill.CreateSkills).Methods("POST")
	router.HandleFunc("/api/skills", skill.GetAllSkills).Methods("GET")

	return router
}
