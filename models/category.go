package models

type Category struct {
	BaseModel
	Name        string     `orm:"size(64);unique" json:"name"`
	Description string     `orm:"size(255)" json:"description"`
	Products    []*Product `orm:"reverse(many)" json:"products"`
}

func (c *Category) TableName() string {
	return "categorias"
}
