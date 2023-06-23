package main

import (
	"bytes"
	"testing"
)

type spySleeper struct{ calls int }

func (s *spySleeper) Sleep() {
	s.calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleep := &spySleeper{}

	Countdown(buffer, spySleep)
	responseBuffer := buffer.String()
	expectBuffer := "3\n2\n1\nGo!\n"

	responseSleepCalls := spySleep.calls
	expectSleepCalls := 3

	if responseSleepCalls != expectSleepCalls {
		t.Errorf("\nSleep error\nexpect:%v\nresponse:%v", expectSleepCalls, responseSleepCalls)

	}

	if responseBuffer != expectBuffer {
		t.Errorf("\nOutput error:\nexpect: %v\nresponse:%v", expectBuffer, responseBuffer)
	}

}
