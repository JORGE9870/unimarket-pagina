package models

type ProductAnalytics struct {
	ViewCount     int     `json:"view_count"`
	SalesCount    int     `json:"sales_count"`
	ReviewsAvg    float64 `json:"reviews_avg"`
	StockTurnover float64 `json:"stock_turnover"`
	TrendScore    float64 `json:"trend_score"`
}
