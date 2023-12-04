package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"gokiwi.com/credit_card_recommender/db"
	"gokiwi.com/credit_card_recommender/handlers"
)

func main() {
	mongodbClient := db.GetMongoDBClient()
	defer mongodbClient.Cleanup()

	// http.HandleFunc("/init", handlers.InitModels)
	http.HandleFunc("/add", handlers.AddCreditCard)
	http.HandleFunc("/addAll", handlers.AddAllCreditCard)
	http.HandleFunc("/all", handlers.GetCreditCards)
	http.HandleFunc("/top", handlers.GetTopCards)

	log.Printf("started server")
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
