package routers

import (
	"unimarket/controllers"

	"unimarket/middleware/auth"
	"unimarket/middleware/cache"
	"unimarket/middleware/cors"
	"unimarket/middleware/error"
	"unimarket/middleware/logging"
	"unimarket/middleware/metrics"
	"unimarket/middleware/ratelimit"
	"unimarket/middleware/validation"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	ns := web.NewNamespace("/v1",
		web.NSBefore(
			cors.HandleCORS,
			logging.LogRequest,
			metrics.CollectMetrics,
			error.HandleError,
		),
		web.NSNamespace("/productos",
			web.NSBefore(
				auth.ValidateToken,
				ratelimit.LimitRequests,
				validation.ValidateProduct,
				cache.CacheProduct,
			),
			web.NSRouter("/",
				&controllers.ProductController{},
				"post:Post;get:GetAll",
			),
			web.NSRouter("/:id",
				&controllers.ProductController{},
				"get:GetOne;put:Put;delete:Delete",
			),
		),
	)
	web.AddNamespace(ns)
}
