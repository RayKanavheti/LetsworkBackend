package main

import (
	"log"
	"net/http"

	"github.com/raykanavheti/LetsworkBackend/models"
	"github.com/raykanavheti/LetsworkBackend/routers"
)

func main() {
	models.InitDB()
	router := routers.InitRoutes()
	log.Fatal(http.ListenAndServe(":8000", router))
}
