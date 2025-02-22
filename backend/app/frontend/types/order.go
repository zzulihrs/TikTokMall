package types

type OrderItem struct {
	ProductName string
	Picture     string
	Qty         int32
	Cost        float32
}

type Order struct {
	OrderId     string
	OrderStatus uint32
	CreatedDate string
	Cost        float32
	Items       []OrderItem
}
