package utils

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// PerClientRateLimiter provides rate limiting per client IP.
func PerClientRateLimiter() gin.HandlerFunc {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	// Background goroutine to clean up inactive clients
	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()
	
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		if _, found := clients[ip]; !found {
			clients[ip] = &client{limiter: rate.NewLimiter(10, 20)}
		}
		clients[ip].lastSeen = time.Now()

		if !clients[ip].limiter.Allow() {
			mu.Unlock()

			// Return 429 Too Many Requests response
			c.JSON(http.StatusTooManyRequests, gin.H{
				"status": "Request Failed",
				"body":   "The API is at capacity, try again later.",
			})
			c.Abort()
			return
		}

		mu.Unlock()
		c.Next()
	}
}
