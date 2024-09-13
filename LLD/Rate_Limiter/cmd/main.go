package main

import (
	"Rate_Limiter/internal"
	"fmt"
	"time"
)

func main() {
	limiter := internal.NewRateLimiter(10*time.Second, 5) // 10-second window, max 5 requests

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		if limiter.AllowRequest() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
	}
}
