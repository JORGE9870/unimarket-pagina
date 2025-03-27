package error

import (
	"github.com/beego/beego/v2/server/web/context"
)

type ErrorMiddleware struct{}

func (m *ErrorMiddleware) HandleError(ctx *context.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Output.JSON(map[string]interface{}{
				"error":   "Error interno del servidor",
				"details": err,
			}, 500, false, false)
		}
	}()

	ctx.Next()

	// Manejar errores específicos
	if err := ctx.Input.GetData("error"); err != nil {
		switch e := err.(type) {
		case *ValidationError:
			ctx.Output.JSON(map[string]interface{}{
				"error":   "Error de validación",
				"details": e.Details,
			}, 400, false, false)
		case *AuthError:
			ctx.Output.JSON(map[string]interface{}{
				"error":   "Error de autenticación",
				"details": e.Message,
			}, 401, false, false)
		default:
			ctx.Output.JSON(map[string]interface{}{
				"error":   "Error inesperado",
				"message": e.Error(),
			}, 500, false, false)
		}
	}
}
