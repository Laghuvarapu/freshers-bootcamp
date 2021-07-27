package Models

type Order struct {
	ID        string   `json:"id"`
	CustomerId   string   `json:"customer_id" `
	Customer Customer `gorm:"foreignKey:CustomerId:"`
	ProductId  string   `json:"product_id" `
	Product Product `gorm:"foreignKey:ProductId:"`
	Quantity  uint      `json:"quantity" `
	Status    string    `json:"status"`
}


func (b *Order) TableName() string {
    return "orders"
}
