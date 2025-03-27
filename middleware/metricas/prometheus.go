package metrics

import (
	"time"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsMiddleware struct {
	requestDuration *prometheus.HistogramVec
	requestTotal    *prometheus.CounterVec
}

func (m *MetricsMiddleware) CollectMetrics(ctx *context.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path
	method := ctx.Request.Method

	// Ejecutar siguiente middleware/controlador
	ctx.Next()

	// Registrar métricas
	duration := time.Since(start).Seconds()
	status := ctx.ResponseWriter.Status

	m.requestDuration.WithLabelValues(path, method, string(status)).Observe(duration)
	m.requestTotal.WithLabelValues(path, method, string(status)).Inc()
}
