package middleware

import (
	"github.com/beego/beego/v2/server/web/context"
)

func LogRequest(ctx *context.Context) {
	// TODO: Implement request logging
	ctx.Output.SetStatus(200)
}
