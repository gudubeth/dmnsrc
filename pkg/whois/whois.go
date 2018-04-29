package whois

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/domainr/whois"
)

// Record represents parsed whois record
type Record struct {
	Name       string
	Response   string
	Error      error
	Elapsed    time.Duration
	Attributes Attributes
}

// Attributes represents attributes of the whois record
type Attributes struct {
	Available bool
	Expires   time.Time
}

// Lookup fetches whois record of specified domain name.
//
// It returns whois data and any error encountered.
func Lookup(name string) (string, error) {
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

// Parse parses whois response. this is a temporary and a very simple
// implementation and quite possibly incorrect
func Parse(name, text string) *Attributes {
	attr := &Attributes{}
	if strings.Contains(text, "No match") ||
		strings.Contains(text, "No entries") ||
		strings.Contains(text, "NOT FOUND") {
		attr.Available = true
	} else {
		attr.Available = false
	}

	return attr
}

// LookupMultiple fetches multiple whois records and puts them to unbuffered channel, which means,
// if you don't read the channel, it will not progress
func LookupMultiple(ctx context.Context, names []string, numWorkers int) <-chan Record {
	if ctx == nil {
		ctx = context.Background()
	}

	//whois library doesn't permit canceling requests. so a request, once started, has to be
	//waited until it is completed. so canceling name generation should be enough for canceling
	//the job
	nc := generateName(ctx, names)
	out := make(chan Record)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for name := range nc {
				rec := fetchForChannel(name)
				out <- *rec
			}
		}()
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

// fetchForChannel lookups for whois record, parses it and prepares Record
// object to send from channel
func fetchForChannel(name string) *Record {
	start := time.Now()
	response, err := Lookup(name)

	rec := &Record{
		Name:     name,
		Response: response,
		Error:    err,
		Elapsed:  time.Since(start),
	}

	if err == nil {
		attr := Parse(name, response)
		rec.Attributes = *attr
	}

	return rec
}
