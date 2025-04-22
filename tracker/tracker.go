package tracker

import (
	"fmt"
	"github.com/fatih/color"
	"runtime"
	"sync"
	"time"
)

var (
	StartColor   = color.New(color.FgGreen).SprintFunc()
	FinishColor  = color.New(color.FgCyan).SprintFunc()
	SummaryColor = color.New(color.FgYellow).SprintFunc()
)

type Tracker struct {
	mu       sync.Mutex
	started  int
	finished int
}

func (t *Tracker) Launch(label string, fn func()) {
	t.mu.Lock()
	t.started++
	logWithTimestamp(StartColor("Starting goroutine"), label)
	t.mu.Unlock()

	go func() {
		defer func() {
			t.mu.Lock()
			t.finished++
			logWithTimestamp(FinishColor("Finished goroutine"), label)
			t.mu.Unlock()
		}()
		fn()
	}()
}

func (t *Tracker) Summary() {
	t.mu.Lock()
	defer t.mu.Unlock()

	now := time.Now().Format("15:04:05")

	fmt.Println()
	fmt.Println(SummaryColor("--- Goroutine Summary ---"), "@", now)
	fmt.Println("Total started: ", t.started)
	fmt.Println("Total finished:", t.finished)
	fmt.Println("Currently running:", runtime.NumGoroutine())
	fmt.Println(SummaryColor("----------------------------"))
	fmt.Println()
}

func logWithTimestamp(prefix string, label string) {
	fmt.Printf("[%s] %s: %s\n",
		time.Now().Format("15:04:05"),
		prefix,
		label,
	)
}
