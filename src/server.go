package main

import (
	"log"
	"net/http"
	"strconv"
	"workshop/src/routes"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	database "workshop/src/database"
)

func main() {

	// Create conts, cant change this
	const PORT = 8080

	// Find type automatically
	appName := "My server"

	// Create slice of strings
	var corsValues []string

	corsValues = append(corsValues, "http://localhost:8080")
	corsValues = append(corsValues, "http://localhost:8000")

	database.Init("psql:test")

	router := httprouter.New()

	router.POST("/user/register", routes.Register)
	router.POST("/benefit", routes.SaveBenefit)

	// Loop over string slice
	for _, value := range corsValues {
		log.Print("Cors value ", value)
	}

	// Create server handle with options
	handler := cors.New(cors.Options{
		AllowedOrigins:   corsValues,
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
	}).Handler(router)

	log.Printf("Starting %s at %d ", appName, PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), handler))
}
