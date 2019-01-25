package main

import (
	"log"
	"net/http"

	"github.com/raykanavheti/LetsworkBackend/models"
	"github.com/raykanavheti/LetsworkBackend/routers"
	"github.com/rs/cors"
)

func main() {
	models.InitDB()
	router := routers.InitRoutes()
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
