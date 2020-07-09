package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// DefaultSleeper counts time
type DefaultSleeper struct{}

// Sleep for DefaultSleeper
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Sleeper implements Sleep()
type Sleeper interface {
	Sleep()
}

// SpySleeper contains Calls
type SpySleeper struct {
	Calls int
}

// Sleep for SpySleeper
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Countdown counts down from 3
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
