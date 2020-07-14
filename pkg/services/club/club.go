package club

// Club represents Club information
type Club struct {
	ID   string `json:"id" firestore:"id"`
	Name string `json:"name" firestore:"name"`
	City string `json:"city" firestore:"city"`
}
