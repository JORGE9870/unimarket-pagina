package models

type Permission struct {
	BaseModel
	Name        string `orm:"size(64);unique" json:"name"`
	Description string `orm:"size(255)" json:"description"`
	CodeName    string `orm:"size(128);unique" json:"code_name"`
}

func (p *Permission) TableName() string {
	return "permisos"
}
