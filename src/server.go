package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"workshop/src/routes"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	database "workshop/src/database"
)

const PORT = 8080

func main() {
	database.Init("psql:test")

	router := httprouter.New()

	router.POST("/user/register", routes.Register)
	router.POST("/benefit", routes.SaveBenefit)

	handler := cors.New(cors.Options{
		AllowedOrigins:   strings.Split("http://localhost:8080,http://localhost:8000", ","),
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
	}).Handler(router)

	log.Print("Starting server at ", PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), handler))
}
