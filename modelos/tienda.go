package models

type Store struct {
	BaseModel
	Name        string     `orm:"size(64);unique" json:"name"`
	Description string     `orm:"size(255)" json:"description"`
	Address     string     `orm:"size(255)" json:"address"`
	Phone       string     `orm:"size(20)" json:"phone"`
	IsActive    bool       `orm:"default(true)" json:"is_active"`
	Owner       *User      `orm:"rel(fk)" json:"owner"`
	Products    []*Product `orm:"reverse(many)" json:"products"`
}

func (s *Store) TableName() string {
	return "tiendas"
}
