package middleware

import (
	"github.com/beego/beego/v2/server/web/context"
)

func LimitRequests(ctx *context.Context) {
	// TODO: Implement rate limiting
	ctx.Output.SetStatus(200)
}
