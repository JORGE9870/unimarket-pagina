package models

type InventoryProjection struct {
	Days     int     `json:"dias"`
	Trend    float64 `json:"tendencia"`
	Stock    int     `json:"stock"`
	MinStock int     `json:"stock_minimo"`
}
