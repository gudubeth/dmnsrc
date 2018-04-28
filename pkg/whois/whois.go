package whois

import (
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
	return response.String(), err
}

type Record struct {
	Name      string
	Response  string
	Error     error
	Elapsed   time.Duration
	Available bool
	Expires   time.Time
}

func FetchMultiple(names []string, numFetchers int) <-chan Record {
	nc := generateName(names)
	out := make(chan Record)
	var wg sync.WaitGroup
	wg.Add(numFetchers)

	for i := 0; i < numFetchers; i++ {
		go func() {
			defer wg.Done()
			for d := range parse(fetch(nc)) {
				out <- d
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func generateName(names []string) <-chan string {
	cn := make(chan string)

	go func() {
		defer close(cn)
		for _, name := range names {
			cn <- name
		}
	}()

	return cn
}

func fetch(cn <-chan string) <-chan Record {
	out := make(chan Record)
	go func() {
		defer close(out)
		for name := range cn {
			start := time.Now()
			repsonse, err := Fetch(name)
			elapsed := time.Since(start)
			out <- Record{
				Name:     name,
				Response: repsonse,
				Error:    err,
				Elapsed:  elapsed,
			}
		}
	}()

	return out
}

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
