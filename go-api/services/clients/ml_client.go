package clients

type MLItem struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type MLClient interface {
	GetItem(id int64) (MLItem, error)
}
