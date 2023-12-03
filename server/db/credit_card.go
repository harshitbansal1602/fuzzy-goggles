package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCreditCards() ([]CreditCardDocument, error) {
	client := GetMongoDBClient()
	coll := client.Client.Database(DB_NAME).Collection(CREDIT_CARD_COLLECTION)
	docs, err := coll.Find(context.TODO(), bson.D{})
	var creditCards []CreditCardDocument

	if err != nil {
		log.Println(err.Error())
		return creditCards, err
	}

	err = docs.All(context.TODO(), &creditCards)

	if err != nil {
		log.Println(err.Error())
		return creditCards, err
	}
	return creditCards, nil
}
