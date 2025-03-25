package models

type Product struct {
	BaseModel
	Name        string    `orm:"size(128)" json:"name"`
	Description string    `orm:"size(255)" json:"description"`
	Price       float64   `orm:"digits(10);decimals(2)" json:"price"`
	Stock       int       `orm:"default(0)" json:"stock"`
	Store       *Store    `orm:"rel(fk)" json:"store"`
	Category    *Category `orm:"rel(fk)" json:"category"`
	IsActive    bool      `orm:"default(true)" json:"is_active"`
}

func (p *Product) TableName() string {
	return "productos"
}
