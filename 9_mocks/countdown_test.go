package main

/*

The requirements for this mock test are to output a countdown and then a final GO!
3
2
1
Go!

We want to test the output functionality, regardless of the string.
Break this down into tasks we get:

- Print 3
- Print 3, 2, 1 and Go!
- Wait a second between each line

*/

import (
	"bytes"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	/*
		t.Run("sleep before every print", func(t *testing.T) {
			spySleepPrinter := &SpyCountdownOperations{}
			Countdown(spySleepPrinter, spySleepPrinter)

			want := []string{
				write,
				sleep,
				write,
				sleep,
				write,
				sleep,
				write,
			}

			if reflect.DeepEqual(want, spySleepPrinter.Calls) {
				t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
			}
		})
	*/

	t.Run("3 prints to Go!", func(t *testing.T) {
		// Recall that bytes.Buffer implements io.Writer
		// In this test, the buffer is sent to io.Writer
		// In main(), the buffer is sent to os.Stdout

		// this creates a pointer to a bytes.Buffer that we'll write to
		buffer := &bytes.Buffer{}

		// here's our mock recorder
		spySleeper := &SpySleeper{}

		// We send both pointers to Countdown
		// one for the buffer, the other for updating the mock
		Countdown(buffer, spySleeper)
		// Countdown(buffer)

		got := buffer.String()
		// first test was just to print '3'
		// once we got that working, we can iterate on what is printed
		// like the multi-line buffer countdown
		// want := "3"
		// this is the multi-line print:
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
		}
	})
}

// Create a struct that "spies" on the interface and counts if it's called
type SpySleeper struct {
	Calls int
}

// This method writes (through the pointer) to the struct
// its receiver is the SpySleeper struct
// so that when the Sleep() method is called on the struct
// it does nothing but increment a counter (it doesn't sleep 1 second)
//
//	note here that in this version, this "Sleep()" is the fake one
//	it doesn't sleep, it just counts monotonically
func (s *SpySleeper) Sleep() {
	// increment the value pointed to by 's' by exactly 1
	s.Calls++
}

// The following creates mocks for testing the order of events
type SpyCountdownOperations struct {
	Calls []string
}

// The version of Sleep() method for the SpyCountdownOperations receiver
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// this is io.Writer being implemented on SpyCountdownOperations
// i.e.: 'Write(p []byte) (n int, err error)' is the signature of Writer
// but instead of writing to anything, it just appends to the Spy and returns
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// here's a type and method for a spy to test configurable sleeper
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	// here's what we want to configure
	sleepTime := 5 * time.Second

	// create a new pointer to spyTime
	spyTime := &SpyTime{}
	// new sleeper struct with the configuration sent
	// remember that spyTime.Sleep here is a mock, it won't actually sleep
	// all it does is record the call, but it has the same signature
	// as a legitimage use of the type
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
