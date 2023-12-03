package handlers

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"slices"
	"sort"

	"gokiwi.com/credit_card_recommender/db"
)

func AddCreditCardsFromCsv(w http.ResponseWriter, r *http.Request) {
	// Specify the path to your CSV file

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func AddCreditCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creditCardRequest NewCreditCard
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&creditCardRequest)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Request malformed", http.StatusBadRequest)
		return
	}

	creditCard := db.CreditCardDocument{
		Name:     creditCardRequest.Name,
		Category: creditCardRequest.Category,
	}

	for _, benefit := range creditCardRequest.Benefits {
		creditCard.Benefits = append(creditCard.Benefits, db.Benefit{
			Category:   benefit.Category.String(),
			Percentage: benefit.Percentage,
			Fixed:      benefit.Fixed,
			Limit:      benefit.Limit,
		})
	}

	client := db.GetMongoDBClient()
	coll := client.Client.Database(db.DB_NAME).Collection(db.CREDIT_CARD_COLLECTION)
	result, err := coll.InsertOne(context.TODO(), creditCard)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to add doc", http.StatusInternalServerError)
		return
	}
	log.Printf("New credit card inserted %v", result.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.InsertedID)
}

func GetCreditCards(w http.ResponseWriter, r *http.Request) {
	creditCards, err := db.GetAllCreditCards()
	if err != nil {
		http.Error(w, "Failed to fetch docs", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creditCards)
}

func GetTopCards(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var topRequest TopCreditCards
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&topRequest)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Request malformed", http.StatusBadRequest)
		return
	}

	if topRequest.Work == 0 {
		http.Error(w, "Work should be present", http.StatusBadRequest)
		return
	}

	if len(topRequest.Lifestyle) == 0 {
		http.Error(w, "Lifestyle length must not be zero", http.StatusBadRequest)
		return
	}

	creditCards, err := db.GetAllCreditCards()
	if err != nil {
		http.Error(w, "Failed to fetch docs", http.StatusInternalServerError)
		return
	}

	distribution := MakeDistribution(topRequest.Lifestyle)

	var creditCardsWithScore []CreditCardsWithScore
	maxScore := float32(0)

	for _, cc := range creditCards {
		score := float32(0)
		for _, benefits := range cc.Benefits {
			category, err := CategoryString(benefits.Category)
			if err != nil {
				log.Println(err.Error())
				log.Printf("Failed to convert following category: %v", benefits.Category)
				continue
			}
			if _, ok := distribution[category]; !ok {
				log.Printf("Missing category in distribution: %v", category)
				continue
			}
			score += float32(math.Min(float64(distribution[category]*benefits.Percentage*WorkIncomeMap[topRequest.Work]/100+benefits.Fixed), float64(benefits.Limit)))
			maxScore = float32(math.Max(float64(maxScore), float64(score)))
		}
		creditCardsWithScore = append(creditCardsWithScore, CreditCardsWithScore{
			CreditCard: cc,
			Score:      score,
			Savings:    score,
		})
	}

	for i := range creditCardsWithScore {
		creditCardsWithScore[i].Score = creditCardsWithScore[i].Score / maxScore * 100
	}

	sort.Slice(creditCardsWithScore, func(i, j int) bool {
		return creditCardsWithScore[i].Score < creditCardsWithScore[j].Score
	})

	slices.Reverse(creditCardsWithScore)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creditCardsWithScore)
}
