package models

import (
	"time"
)

type BaseModel struct {
	Id        int64     `orm:"auto;pk" json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
} 