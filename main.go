package main

import (
	"fmt"
	"time"

	"track-a-rou/tracker"
)

func main() {
	t := &tracker.Tracker{}

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			t.Summary()
		}
	}()

	// Simulate launching 10 goroutines
	for i := 0; i < 10; i++ {
		i := i
		t.Launch(
			workerLabel(i),
			func() {
				time.Sleep(time.Duration(1+i%3) * time.Second)
			},
		)
		time.Sleep(300 * time.Millisecond)
	}

	time.Sleep(10 * time.Second)
}

func workerLabel(i int) string {
	return "worker-" + fmt.Sprint(i)
}
