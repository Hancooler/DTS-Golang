package item

// Service defines the interface for interacting with items
type Service interface {
	Get() ([]Item, error)
	Create(input CreateItemInput) (Item, error) // Use a clear input struct name
}

// CreateItemInput defines the structure for creating a new item
type CreateItemInput struct {
	ItemCode    string `json:"item_code"`   // Add JSON tag for potential marshalling
	Description string `json:"description"` // Add JSON tag
	Quantity    int    `json:"quantity"`    // Add JSON tag
	OrderID     int64  `json:"order_id"`    // Assuming OrderID can be larger than int
}

type service struct {
	repository Repository
}

// NewService creates a new service instance
func NewService(repository Repository) *service {
	return &service{repository: repository}
}

// Get retrieves all items from the underlying repository
func (s *service) Get() ([]Item, error) {
	items, err := s.repository.Get()
	if err != nil {
		return nil, err // Return nil for items and propagate the error
	}
	return items, nil
}

// Create creates a new item using the repository
func (s *service) Create(input CreateItemInput) (Item, error) {
	// No need to define a nested Item struct here, use the existing Item type
	itemCreated, err := s.repository.Create(Item{
		ItemCode:    input.ItemCode,
		Description: input.Description,
		Quantity:    input.Quantity,
		OrderID:     input.OrderID,
	})
	if err != nil {
		return Item{}, err // Return a default Item and propagate the error
	}
	return itemCreated, nil
}
