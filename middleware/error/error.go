package error

import (
	"github.com/beego/beego/v2/server/web/context"
)

func HandleError(ctx *context.Context) {
	// TODO: Implement error handling
	ctx.Output.SetStatus(200)
}
