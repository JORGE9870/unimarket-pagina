package controladores

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type ControladorBase struct {
	web.Controller
}

type Respuesta struct {
	Exito bool        `json:"exito"`
	Datos interface{} `json:"datos,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (c *ControladorBase) ParsearYValidarJSON(v interface{}) error {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, v); err != nil {
		return err
	}
	return nil
}

func (c *ControladorBase) RespuestaExito(datos interface{}) {
	c.Data["json"] = Respuesta{
		Exito: true,
		Datos: datos,
	}
	c.ServeJSON()
}

func (c *ControladorBase) RespuestaError(err string, code int) {
	c.Ctx.Output.SetStatus(code)
	c.Data["json"] = Respuesta{
		Exito: false,
		Error: err,
	}
	c.ServeJSON()
}

func (c *ControladorBase) ObtenerIDUsuario() int64 {
	return c.Ctx.Input.GetData("user_id").(int64)
}

func (c *ControladorBase) ObtenerRolesUsuario() []string {
	return c.Ctx.Input.GetData("roles").([]string)
}

func (c *ControladorBase) TienePermiso(permiso string) bool {
	roles := c.ObtenerRolesUsuario()
	return contieneAlguno(roles, []string{"admin", permiso})
}

func contieneAlguno(slice []string, valores []string) bool {
	for _, v := range valores {
		for _, s := range slice {
			if s == v {
				return true
			}
		}
	}
	return false
}
