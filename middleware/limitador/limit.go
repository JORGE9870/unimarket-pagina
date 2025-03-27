package ratelimit

import (
	"sync"

	"github.com/beego/beego/v2/server/web/context"
	"golang.org/x/time/rate"
)

type RateLimiter struct {
	visitors map[string]*rate.Limiter
	mu       sync.Mutex
	rate     rate.Limit
	burst    int
}

func (rl *RateLimiter) LimitRequests(ctx *context.Context) {
	ip := ctx.Input.IP()

	rl.mu.Lock()
	limiter, exists := rl.visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(rl.rate, rl.burst)
		rl.visitors[ip] = limiter
	}
	rl.mu.Unlock()

	if !limiter.Allow() {
		ctx.Output.JSON(map[string]interface{}{
			"error": "Demasiadas solicitudes",
		}, 429, false, false)
		return
	}
}
