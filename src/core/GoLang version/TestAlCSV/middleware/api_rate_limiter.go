package middleware

import (

	// Third-party libraries
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	// System Packages
	"net/http"
	"sync"
)

// RateLimiterMiddleware limits the number of requests for each IP
func RateLimiterMiddleware(limit rate.Limit, burst int) gin.HandlerFunc {
	// Creating a card to store limiters for each IP
	visitors := make(map[string]*rate.Limiter)
	var mu sync.Mutex

	// A function for obtaining or creating an IP limiter
	getLimiter := func(ip string) *rate.Limiter {
		mu.Lock()
		defer mu.Unlock()

		limiter, exists := visitors[ip]
		if !exists {
			// Creating a new limiter with the specified limit and the size of the "bucket"
			limiter = rate.NewLimiter(limit, burst)
			visitors[ip] = limiter
		}
		return limiter
	}

	return func(c *gin.Context) {
		ip := c.ClientIP() // Getting the client's IP address
		limiter := getLimiter(ip)

		if !limiter.Allow() {
			// If the limit is exceeded, we return an error.
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests",
			})
			return
		}

		c.Next()
	}
}
