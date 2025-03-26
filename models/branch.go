package models

type Branch struct {
	BaseModel
	Name     string `orm:"size(64)" json:"name"`
	Address  string `orm:"size(255)" json:"address"`
	Phone    string `orm:"size(20)" json:"phone"`
	IsActive bool   `orm:"default(true)" json:"is_active"`
	Store    *Store `orm:"rel(fk)" json:"store"`
}

func (b *Branch) TableName() string {
	return "sucursales"
}
