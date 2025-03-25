package models

type Role struct {
	BaseModel
	Name        string        `orm:"size(64);unique" json:"name"`
	Description string        `orm:"size(255)" json:"description"`
	Permissions []*Permission `orm:"rel(m2m)" json:"permissions"`
}

func (r *Role) TableName() string {
	return "roles"
}
