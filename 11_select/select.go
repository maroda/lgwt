package main

import (
	"fmt"
	"net/http"
	"time"
)

// to be used 'in production' with Racer()
var tenSecondTimeout = 10 * time.Second

// this is the 'production' function that uses the configurable one underneath
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// to test, this configurable one can be used directly so it doesn't have to be 10s
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// select is allowing us to wait on multiple channels,
	// sending them all at once
	// whichever writes to its channel first is returned
	// each 'case' is a variable waiting for input on the channel
	// which it gets as a return value from ping()
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
		// time.After measures the time since the select began,
		// and returns a struct channel just like the others.
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

	/* ... without a select statement
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	/* ... without a function ...
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
	*/
}

// pulling out the function that does the getting
/* ... without a select statement, this isn't needed
func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
*/

// this returns a struct{} type because
// there's no data being returned, and
// a struct uses the least amount of mem
func ping(url string) chan struct{} {
	// make the channel
	ch := make(chan struct{})
	// when the http.Get is complete, close the associated channel
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
