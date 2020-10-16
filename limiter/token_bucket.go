package limiter

import (
	"golang.org/x/time/rate"
	"time"
)

func NewTokenBucket(n int, b int) *rate.Limiter {
	l := rate.Every(time.Second / time.Duration(n))
	lim := rate.NewLimiter(l, b)
	return lim
}
