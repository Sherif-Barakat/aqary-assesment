package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
	"golang.org/x/time/rate"
)

type ip_address string

type rate_limiter struct {
	ips map[ip_address]*rate.Limiter
	mu  *sync.RWMutex
}

func new_rl() *rate_limiter {
	return &rate_limiter{
		ips: make(map[ip_address]*rate.Limiter),
		mu:  &sync.RWMutex{},
	}
}

func (rl *rate_limiter) getLimiter(ip ip_address) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(1), 5) // Allow 1 request per second with a burst of 5
		rl.ips[ip] = limiter
	}

	return limiter
}

func (rl *rate_limiter) handler(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	limiter := rl.getLimiter(ip_address(ip))

	if !limiter.Allow() {
		http.Error(w, "Too many requests, please try again later.", http.StatusTooManyRequests)
		return
	}

	w.Write([]byte("Data received successfully!"))
}

func main() {
	rl := new_rl()

	http.HandleFunc("/api/data", rl.handler)

	port := ":8000"
	fmt.Println("Server listening on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
