package whois

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/domainr/whois"
)

// Fetch fetches whois record of specified domain name.
//
// It returns whois data and any error encountered.
func Fetch(name string) (string, error) {

	request, err := whois.NewRequest(name)
	if err != nil {
		return "", err
	}
	var client = whois.NewClient(whois.DefaultTimeout)
	response, err := client.Fetch(request)
	if err != nil {
		return "", err
	}
	return response.String(), err
}

// Record presents parsed whois record
type Record struct {
	Name      string
	Response  string
	Error     error
	Elapsed   time.Duration
	Available bool
	Expires   time.Time
}

// FetchMultiple fetches multiple whois records and puts them to unbuffered channel, which means,
// if you don't read the channel, it will not progress
func FetchMultiple(ctx context.Context, names []string, numFetchers int) <-chan Record {
	if ctx == nil {
		ctx = context.Background()
	}

	//whois library doesn't permit canceling requests. so a request, once started, has to be
	//waited until it is completed. so canceling name generation should be enough for canceling
	//the job
	nc := generateName(ctx, names)
	out := make(chan Record)
	var wg sync.WaitGroup
	wg.Add(numFetchers)

	for i := 0; i < numFetchers; i++ {
		go func(i int) {
			defer wg.Done()

			//TODO remove pipeline, this style is not really needed here
			for rec := range parse(fetch(nc)) {
				out <- rec
			}

		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// generateName converts names array to unbuffered channel
func generateName(ctx context.Context, names []string) <-chan string {
	nc := make(chan string)

	go func() {
		defer close(nc)
		for _, name := range names {
			select {
			case nc <- name:
			case <-ctx.Done():
				return
			}
		}
	}()

	return nc
}

// fetch fetches whois records coming from cn
func fetch(cn <-chan string) <-chan Record {
	out := make(chan Record)
	go func() {
		defer close(out)
		for name := range cn {
			start := time.Now()
			repsonse, err := Fetch(name)
			out <- Record{
				Name:     name,
				Response: repsonse,
				Error:    err,
				Elapsed:  time.Since(start),
			}
		}
	}()

	return out
}

// parse parses whois text. this is a temporary and a very simple implementation,
// quite possibly incorrect
func parse(recs <-chan Record) <-chan Record {
	out := make(chan Record)
	go func() {
		defer close(out)
		for r := range recs {
			if r.Error == nil &&
				(strings.Contains(r.Response, "No match") ||
					strings.Contains(r.Response, "No entries") ||
					strings.Contains(r.Response, "NOT FOUND")) {
				r.Available = true
			} else {
				r.Available = false
			}
			out <- r
		}
	}()
	return out
}
