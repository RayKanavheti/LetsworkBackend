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
	return router
}
