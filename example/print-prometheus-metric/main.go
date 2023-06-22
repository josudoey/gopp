package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/common/expfmt"
)

func main() {
	r := prometheus.NewRegistry()
	requestsTotal := promauto.With(r).NewCounterVec(prometheus.CounterOpts{
		Name: "requests_total",
	}, []string{"instance"})

	counter := requestsTotal.WithLabelValues(fmt.Sprintf("joey-%d", os.Getpid()))

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			counter.Inc()
		}
	}()

	buf := &bytes.Buffer{}

	for range time.NewTicker(time.Second).C {
		buf.Reset()
		mfs, err := r.Gather()
		if err != nil {
			log.Fatal(err)
		}

		enc := expfmt.NewEncoder(buf, expfmt.FmtText)
		for _, mf := range mfs {
			if err := enc.Encode(mf); err != nil {
				log.Fatal(err)
			}
		}

		fmt.Printf("%v\n", buf.String())
	}

	// Output
	// # HELP requests_total
	// # TYPE requests_total counter
	// requests_total{instance="joey-4508"} 1
	//
	// # HELP requests_total
	// # TYPE requests_total counter
	// requests_total{instance="joey-4508"} 2
	//
}
