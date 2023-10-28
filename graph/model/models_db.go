package model

type TrainingDB struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Coast    int    `json:"coast"`
	GymID    string `json:"gym_id"`
}

type CustomerDB struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GymDB struct {
	ID     string `json:"id"`
	Branch string `json:"branch"`
	Admin  string `json:"admin"`
	Phone  string `json:"phone"`
	Slots  int    `json:"slots"`
}

type PurchaseDB struct {
	ID         string  `json:"id"`
	TrainingID string  `json:"training_id"`
	CustomerID string  `json:"customer_id"`
	Coast      int     `json:"coast"`
	Income     float64 `json:"income"`
}
