package routers

import (
	"github.com/gorilla/mux"
	"github.com/raykanavheti/LetsworkBackend/controllers"
)

var (
	user = new(controllers.UserController)
)

//InitRoutes : Registering all system Routes.
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/user", user.CreateUser).Methods("POST")
	router.HandleFunc("/api/users", user.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", user.GetUserByID).Methods("GET")
	router.HandleFunc("/api/user_confirm/{uuid}", user.GetUserByUUID).Methods("GET")
  router.HandleFunc("/api/reset_password/{email}", user.SendRestLink).Methods("GET")
	return router
}
