package models

type Notification struct {
	BaseModel
	UserId     int64  `orm:"column(id_usuario)" json:"id_usuario"`
	Title      string `orm:"column(titulo)" json:"titulo"`
	Message    string `orm:"column(mensaje)" json:"mensaje"`
	Type       string `orm:"column(tipo)" json:"tipo"`
	IsRead     bool   `orm:"column(leida)" json:"leida"`
	CreateDate string `orm:"column(fecha_creacion)" json:"fecha_creacion"`
}

func (n *Notification) TableName() string {
	return "notificaciones"
}
