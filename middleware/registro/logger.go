package logging

import (
	"time"

	"github.com/beego/beego/v2/server/web/context"
	"go.uber.org/zap"
)

type LoggerMiddleware struct {
	logger *zap.Logger
}

func (m *LoggerMiddleware) LogRequest(ctx *context.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path

	// Ejecutar siguiente middleware/controlador
	ctx.Next()

	// Log después de la respuesta
	m.logger.Info("request",
		zap.String("path", path),
		zap.String("method", ctx.Request.Method),
		zap.Int("status", ctx.ResponseWriter.Status),
		zap.Duration("latency", time.Since(start)),
		zap.String("ip", ctx.Input.IP()),
	)
}
