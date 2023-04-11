package clients

// MockClient is a particular implementation of MLClient (interface)
// And it works returning hardcoded values
// This client doesn't need be tested :)
type MockClient struct{}

func (mockClient MockClient) GetItem(id int64) (MLItem, error) {
	return MLItem{
		Title: "Mocked item title",
		Price: 50.00,
	}, nil
}
