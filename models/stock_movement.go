package models

type StockMovement struct {
	ID        int64  `json:"id"`
	ProductID int64  `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Type      string `json:"type"`
	Date      string `json:"date"`
}
