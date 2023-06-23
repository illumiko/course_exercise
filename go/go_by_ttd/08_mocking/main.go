package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleep interface {
	Sleep()
}
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const countdownEnd = "Go!"

func Countdown(buffer io.Writer, sleeper Sleep) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(buffer, i)
	}
	fmt.Fprintln(buffer, countdownEnd)

}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
