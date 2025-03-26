package controllers

import (
    "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
    web.Controller
}

func (c *BaseController) ResponseSuccess(data interface{}) {
    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data":    data,
    }
    c.ServeJSON()
}

func (c *BaseController) ResponseError(message string, code int) {
    c.Ctx.Output.SetStatus(code)
    c.Data["json"] = map[string]interface{}{
        "success": false,
        "message": message,
    }
    c.ServeJSON()
}