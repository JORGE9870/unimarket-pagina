package routers

import (
	"unimarket/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/usuarios",
			web.NSRouter("/", &controllers.UserController{}, "post:Post;get:GetAll"),
			web.NSRouter("/:id", &controllers.UserController{}, "get:Get;put:Put;delete:Delete"),
		),
		web.NSNamespace("/stores",
			web.NSRouter("/", &controllers.StoreController{}, "post:Create;get:GetAll"),
			web.NSRouter("/:id", &controllers.StoreController{}, "get:Get;put:Update;delete:Delete"),
		),
		web.NSNamespace("/productos",
			web.NSRouter("/", &controllers.ProductController{}, "post:Create;get:GetAll"),
			web.NSRouter("/:id", &controllers.ProductController{}, "get:Get;put:Update;delete:Delete"),
		),
		web.NSNamespace("/categorias",
			web.NSRouter("/", &controllers.CategoryController{}, "post:Create;get:GetAll"),
			web.NSRouter("/:id", &controllers.CategoryController{}, "get:Get;put:Update;delete:Delete"),
		),
		web.NSNamespace("/roles",
			web.NSRouter("/", &controllers.RoleController{}, "post:Create;get:GetAll"),
			web.NSRouter("/:id", &controllers.RoleController{}, "get:Get;put:Update;delete:Delete"),
		),
		web.NSNamespace("/permisos",
			web.NSRouter("/", &controllers.PermissionController{}, "post:Create;get:GetAll"),
			web.NSRouter("/:id", &controllers.PermissionController{}, "get:Get;put:Update;delete:Delete"),
		),
		web.NSNamespace("/pedidos",
			web.NSRouter("/", &controllers.OrderController{}, "post:Create;get:GetAll"),
			web.NSRouter("/:id", &controllers.OrderController{}, "get:Get;put:Update;delete:Delete"),
		),
	)
	web.AddNamespace(ns)
}
