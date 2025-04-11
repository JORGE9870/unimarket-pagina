package cors

import (
	"github.com/beego/beego/v2/server/web/context"
)

func HandleCORS(ctx *context.Context) {
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
	ctx.Output.SetStatus(200)
}
