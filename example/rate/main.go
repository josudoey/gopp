package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	burst := 10
	limter := rate.NewLimiter(3, burst)
	// limter
	for i := 0; i < 3; i++ {
		go func(n int) {
			for {
				limter.Wait(ctx)
				log.Printf("n:%v current limit:%v Burst:%v", n, limter.Limit(), limter.Burst())
			}
		}(i)
	}

	go func() {
		<-time.After(time.Second * 3)
		limter.SetLimit(1)
		<-time.After(time.Second * 5)
		cancel()
	}()

	<-ctx.Done()
	log.Printf("done")
}
