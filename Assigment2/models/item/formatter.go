package item

type ItemFormatted struct {
	ID          int    `json:"id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
}

func ItemFormatter(item Item) ItemFormatted {
	type Item struct {
		ItemCode    string
		Quantity    int
		Description string
		OrderId     int // Assuming OrderId is of type int
	}
	itemFormatted := ItemFormatted{}
	return itemFormatted
}

func ItemsFormatter(items []Item) []ItemFormatted {
	var itemsFormatted []ItemFormatted
	for _, item := range items {
		itemsFormatted = append(itemsFormatted, ItemFormatter(item))
	}

	return itemsFormatted
}
