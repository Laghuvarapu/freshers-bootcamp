package Models

type Customer struct {
   ID        string   `json:"id"`
   CustomerName   string   `json:"customer_name" `
   MpbileNumber  uint   `json:"mobile_number" `
}

func (b *Customer) TableName() string {
      return "customers"
}
