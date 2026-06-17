package domain

// StockItemInput representa os dados que o frontend (ou Postman) vai enviar
type StockItemInput struct {
	SerialNumber string `json:"serial_number"`
	Condition    string `json:"condition"`
}

// StockItemOutput representa os dados completos que a API vai devolver para o Frontend
type StockItemOutput struct {
	SerialNumber string `json:"serial_number"`
	Condition    string `json:"condition"`
	Status       string `json:"status"`
	ProductName  string `json:"product_name"`
	ProductModel string `json:"product_model"`
}
