package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) rateLimit(next http.Handler) http.Handler {

	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var mutex sync.Mutex
	var clients = make(map[string]*client)

	// launch a background goroutine which remove old entries from the clients map once every minute
	go func() {
		for {
			time.Sleep(1 * time.Minute)

			mutex.Lock()

			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}

			mutex.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only carry out the check if rate limiting is enabled
		if app.config.limiter.enable {
			// Extract the client IP address from the request
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}

			mutex.Lock()

			// Check if the ip is in the map and if it's not create a new limiter with the default values
			if _, found := clients[ip]; !found {
				clients[ip] = &client{
					limiter:  rate.NewLimiter(rate.Limit(app.config.limiter.rps), app.config.limiter.burst),
					lastSeen: time.Now(),
				}
			}

			if !clients[ip].limiter.Allow() {
				mutex.Unlock()
				app.rateLimitExceededResponse(w, r)
				return
			}

			mutex.Unlock()
		}

		next.ServeHTTP(w, r)
	})
}
