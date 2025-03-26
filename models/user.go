package models

type User struct {
	BaseModel
	Username  string  `orm:"size(128);unique" json:"username"`
	Email     string  `orm:"size(128);unique" json:"email"`
	Password  string  `orm:"size(128)" json:"-"`
	FirstName string  `orm:"size(128)" json:"first_name"`
	LastName  string  `orm:"size(128)" json:"last_name"`
	IsActive  bool    `orm:"default(true)" json:"is_active"`
	Roles     []*Role `orm:"rel(m2m)" json:"roles"`
}

func (u *User) TableName() string {
	return "usuarios"
}
