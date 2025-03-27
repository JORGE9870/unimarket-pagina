package models

type Tienda struct {
	Id          int64  `json:"id_tienda"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Direccion   string `json:"direccion"`
	CreateDate  string `json:"fecha_creacion"`
}
