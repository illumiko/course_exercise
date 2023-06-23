package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("Testing Counter", func(t *testing.T) {
		c := &Counter{}
		expectedCounter := 1000

		var wg sync.WaitGroup
		wg.Add(expectedCounter)

		for i := 0; i < expectedCounter; i++ {
			go func() {
				c.Increment()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, expectedCounter, c)
	})

}

func assertCounter(t testing.TB, expect int, response *Counter) {
	t.Helper()
	if expect != response.count {
		t.Errorf("\nExpect %v\nResponse %v\n", expect, response)
	}

}
