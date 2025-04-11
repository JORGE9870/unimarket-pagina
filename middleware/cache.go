package middleware

import (
	"github.com/beego/beego/v2/server/web/context"
)

func CacheProduct(ctx *context.Context) {
	// TODO: Implement product caching
	ctx.Output.SetStatus(200)
}
