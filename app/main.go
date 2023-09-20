package main

import (
	"fmt"
	"log"
	"os"
	"technical-test/priverion/routes"
	"technical-test/priverion/utils"

	"github.com/joho/godotenv"
)

func main() {
	// load de envioronments variables
	err := godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// import the set up of mongo
	db, errDb := utils.NewDatabase()
	if errDb != nil {
		log.Fatal(errDb)
	}
	defer db.Close()

	// import the set up router to run the server and import all routes
	router := routes.SetupRouter(db)
	if err != nil {
		fmt.Println("Error to load variables")
	}

	router.Run(fmt.Sprintf("localhost:%s", port))
}
