package models

type Order struct {
	BaseModel
	User        *User        `orm:"rel(fk)" json:"user"`
	Store       *Store       `orm:"rel(fk)" json:"store"`
	Status      string       `orm:"size(32)" json:"status"`
	TotalAmount float64      `orm:"digits(10);decimals(2)" json:"total_amount"`
	Items       []*OrderItem `orm:"reverse(many)" json:"items"`
	Delivery    *Delivery    `orm:"rel(fk)" json:"delivery"`
}

type OrderItem struct {
	BaseModel
	Order      *Order   `orm:"rel(fk)" json:"order"`
	Product    *Product `orm:"rel(fk)" json:"product"`
	Quantity   int      `json:"quantity"`
	UnitPrice  float64  `orm:"digits(10);decimals(2)" json:"unit_price"`
	TotalPrice float64  `orm:"digits(10);decimals(2)" json:"total_price"`
}

type Delivery struct {
	BaseModel
	Order       *Order `orm:"rel(fk)" json:"order"`
	Address     string `orm:"size(255)" json:"address"`
	Status      string `orm:"size(32)" json:"status"`
	DeliveryMan *User  `orm:"rel(fk)" json:"delivery_man"`
}

func (o *Order) TableName() string {
	return "pedidos"
}

func (oi *OrderItem) TableName() string {
	return "detalle_pedidos"
}

func (d *Delivery) TableName() string {
	return "entregas"
}
