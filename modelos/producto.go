package models

type Product struct {
	BaseModel
	Name        string  `orm:"size(128)" json:"name"`
	Description string  `orm:"size(512)" json:"description"`
	Price       float64 `orm:"digits(10);decimals(2)" json:"price"`
	Stock       int     `json:"stock"`
	Store       *Store  `orm:"rel(fk)" json:"store"`
	Category    string  `orm:"size(64)" json:"category"`
	Rating      float64 `json:"rating"`
}

func (p *Product) TableName() string {
	return "productos"
}

type ProductRequest struct {
	Name       string   `json:"name"`
	SKU        string   `json:"sku"`
	BasePrice  float64  `json:"base_price"`
	Discount   float64  `json:"discount"`
	Quantity   int      `json:"quantity"`
	Categories []string `json:"categories"`
	Brand      string   `json:"brand"`
	Images     []string `json:"images"`
}

type ProductResponse struct {
	Name       string   `json:"name"`
	FinalPrice float64  `json:"final_price"`
	Available  bool     `json:"available"`
	Locations  []string `json:"locations"`
}

type ProductAnalytics struct {
	ViewCount     int     `json:"view_count"`
	SalesCount    int     `json:"sales_count"`
	ReviewsAvg    float64 `json:"reviews_avg"`
	StockTurnover float64 `json:"stock_turnover"`
	TrendScore    float64 `json:"trend_score"`
}

type ProductDetails struct {
	ID            int64      `json:"id"`
	Name          string     `json:"nombre"`
	Description   string     `json:"descripcion"`
	Price         float64    `json:"precio"`
	CategoryName  string     `json:"categoria_nombre"`
	BrandName     string     `json:"marca_nombre"`
	AverageRating float64    `json:"rating_promedio"`
	TotalReviews  int64      `json:"total_reviews"`
	TotalSales    int64      `json:"total_ventas"`
	Locations     []Location `json:"ubicaciones"`
}

type Location struct {
	ID       int64  `json:"id"`
	Name     string `json:"nombre"`
	Quantity int    `json:"cantidad"`
}
