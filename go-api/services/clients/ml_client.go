package clients

type MLItem struct {
	Title    string    `json:"title"`
	Price    float64   `json:"price"`
	Pictures []MLImage `json:"pictures"`
}

type MLImage struct {
	URL string `json:"url"`
}

// MLClient is the interface definition for a ML Client
// It only defines how the contract should be
// Then, each specific structure implements the corresponding logic
type MLClient interface {
	GetItem(id int64) (MLItem, error)
}
