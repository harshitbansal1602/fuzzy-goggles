package db

type DB interface {
	init()
	Cleanup()
}

// Document Types
type CreditCardDocument struct {
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Benefits []Benefit `json:"benefits"`
}

type Benefit struct {
	Category    string  `json:"category"`
	Percentage  float32 `json:"percentage"`
	Fixed       float32 `json:"fixed"`
	Limit       float32 `json:"limit"`
	Description string  `json:"description"`
}

const DB_NAME = "kiwi"
const CREDIT_CARD_COLLECTION = "credit_cards"
const TEST_CREDIT_CARD_COLLECTION = "test_credit_cards"
const DRAFT_CREDIT_CARD_COLLECTION = "draft_credit_cards"
