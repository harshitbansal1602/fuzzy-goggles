package handlers

import "gokiwi.com/credit_card_recommender/db"

type CreditCardsWithScore struct {
	CreditCard db.CreditCardDocument `json:"creditCard"`
	Score      float32               `json:"score"`
	Savings    float32               `json:"savings"`
}
