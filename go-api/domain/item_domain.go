package domain

type Item struct {
	ID        int64       `json:"id"`
	Name      string      `json:"name"`
	Price     float64     `json:"price"`
	Thumbnail string      `json:"thumbnail"`
	Pictures  interface{} `json:"pictures"`
}
