package handlers

type Category int

const (
	All Category = iota
	OnlineShopping
	Fuel
	Travel
	Food
	Membership
)

type Lifestyle int

const (
	LongDrives Lifestyle = iota + 1
	Foodie
	ImpulsiveBuyer
	BingeWatcher
	LoveGadgets
)

var LifestyleCategoryMap = map[Lifestyle][]Category{
	LongDrives:     {Travel, Fuel},
	Foodie:         {Food},
	ImpulsiveBuyer: {OnlineShopping},
	BingeWatcher:   {Membership},
	LoveGadgets:    {OnlineShopping},
}

func MakeDistribution(lifestyle []Lifestyle) map[Category]float32 {
	base := 0.3
	boost := 0.05
	distribution := map[Category]float32{All: float32(base)}
	for _, ls := range lifestyle {
		for _, cat := range LifestyleCategoryMap[ls] {
			distribution[cat] += float32(boost)
		}
	}

	total := float32(0)
	for _, percentage := range distribution {
		total += percentage
	}
	scale := 1 / total
	for cat, percentage := range distribution {
		distribution[cat] = percentage * scale
	}
	return distribution
}

type Work int

const (
	Student Work = iota + 1
	JrEmployee
	SrEmployee
	CXO
)

var WorkIncomeMap = map[Work]float32{
	Student:    60000,
	JrEmployee: 500000,
	SrEmployee: 1200000,
	CXO:        2400000,
}

type NewCreditCard struct {
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Benefits []Benefit `json:"benefits"`
}

type Benefit struct {
	Category    Category `json:"category"`
	Percentage  float32  `json:"percentage"`
	Fixed       float32  `json:"fixed"`
	Limit       float32  `json:"limit"`
	Description string   `json:"description"`
}

type Distribution struct {
	Category   Category `json:"category"`
	Percentage float32  `json:"percentage"`
}

type TopCreditCards struct {
	Work      Work        `json:"work"`
	Lifestyle []Lifestyle `json:"lifestyle"`
}
