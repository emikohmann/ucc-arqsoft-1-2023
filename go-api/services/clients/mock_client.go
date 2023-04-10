package clients

type MockClient struct{}

func (mockClient MockClient) GetItem(id int64) (MLItem, error) {
	return MLItem{
		Title: "Mocked item title",
		Price: 50.00,
	}, nil
}
