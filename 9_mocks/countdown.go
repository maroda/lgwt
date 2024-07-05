package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
	write          = "write"
	sleep          = "sleep"
)

// Create an interface that will implement the existing Sleep() method
// We want to use this so we can mock the sleep dependency
// and not have to actually sleep three seconds to test it every time.
//
//	This not used by the "main" code, because that uses a struct
//	that ties Sleep() onto it as a method, which does real sleeping.
type Sleeper interface {
	Sleep()
}

// This is the real sleeper, just a struct, that will have a method
type DefaultSleeper struct{}

// this can be configured to contain a duration and a sleep function
// sleep has the same signature as time.Sleep, so they will match as types
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	// remember that ConfigurableSleeper is TWO THINGS:
	// a duration and a sleep
	// so these get the data from each of those values
	c.sleep(c.duration)
}

// when the struct is implemented, its Sleep method will run
//
//	this is the "real" sleep that uses the real sleep interface
//	DefaultSleeper is the receiver of the Sleep() method
//	and within it we can use the real built-in Sleep functionality
func (d *DefaultSleeper) Sleep() {
	// call sleep in much the same way we were called
	// using a Sleep() method on time to calculate a second of sleep
	time.Sleep(1 * time.Second)
}

// Once we know bytes.Buffer works,
// we need that to be a general interface
// so io.Writer - which is implemented by bytes.Buffer and Fprint - can be used
//
//	i.e.: tests will pass with either version!
//
// func Countdown(out *bytes.Buffer) {
// func Countdown(out io.Writer) {
//
// to test with a mock that will track the testing,
// we need to take a Sleeper interface
func Countdown(out io.Writer, sleeper Sleeper) {
	// Fprint takes an io.Writer and sends a string to it.
	// i.e.: write '3' into the buffer that 'out' points to
	// fmt.Fprint(out, "3")

	//for i := countdownStart; i > 0; i-- {
	//	sleeper.Sleep()
	//}

	// to do a multiline readout, use a for loop and print a line
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		// adding a sleep is a dependency we need to track with a test
		// so it's time to implement it as an interface
		// time.Sleep(1 * time.Second)

		// instead of the direct sleep command, we can now use the interface
		// because sleeper is a pointer to a DefaultSleeper type
		// that has the method on it
		// so it runs when called
		sleeper.Sleep()
	}

	// the last line by itself
	fmt.Fprint(out, finalWord)
}

func main() {
	// create a sleeper
	// it is a pointer to a struct that has a Sleep() method
	// when it gets created, it calls the method, and runs the sleep
	// sleeper := &DefaultSleeper{}
	// or use the new ConfigurableSleeper
	// which requires the duration and a sleep method, of which time.Sleep is one
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
