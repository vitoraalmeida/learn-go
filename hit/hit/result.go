package hit

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Result is a request's result.
type Result struct {
	RPS      float64       // RPS is the requests per second
	Requests int           // Requests is the number of requests made
	Errors   int           // Errors is the number of errors occurred
	Bytes    int64         // Bytes is the number of bytes downloaded
	Duration time.Duration // Duration is a single or all requests' duration
	Fastest  time.Duration // Fastest request result duration among others
	Slowest  time.Duration // Slowest request result duration among others
	Status   int           // Status is a request's HTTP status code
	Error    error         // Error is not nil if the request is failed
}

// Faz a união entre um result e este
func (r *Result) Merge(o *Result) {
	r.Requests++
	r.Bytes += o.Bytes
	if r.Fastest == 0 || o.Duration < r.Fastest {
		r.Fastest = o.Duration
	}
	if o.Duration > r.Slowest {
		r.Slowest = o.Duration
	}

	switch {
	case o.Error != nil:
		fallthrough // executa a instrução do próximo case sem checar sua condição
	case o.Status >= http.StatusBadRequest:
		r.Errors++
	}
}

// Consolida a duração total e a taxa de requests por segundo
func (r *Result) Finalize(total time.Duration) *Result {
	r.Duration = total
	r.RPS = float64(r.Requests) / total.Seconds()
	return r
}

func (r *Result) Fprint(out io.Writer) {
	p := func(format string, args ...any) {
		fmt.Fprintf(out, format, args...)
	}
	p("\nSummary:\n")
	p("\tSuccess	: %.0f%%\n", r.success())
	p("\tRPS		: %.1f\n", r.RPS)
	p("\tRequests	: %d\n", r.Requests)
	p("\tErrors		: %d\n", r.Errors)
	p("\tBytes		: %d\n", r.Bytes)
	p("\tDuration	: %s\n", round(r.Duration))
	if r.Requests > 1 {
		p("\tFastest	: %s\n", round(r.Fastest))
		p("\tSlowest	: %s\n", round(r.Slowest))
	}
}

func (r *Result) success() float64 {
	rr, e := float64(r.Requests), float64(r.Errors)
	return (rr - e) / rr * 100
}

func round(t time.Duration) time.Duration {
	return t.Round(time.Microsecond)
}
