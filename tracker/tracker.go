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

type GoroutineInfo struct {
	Label    string
	Started  time.Time
	Duration time.Duration
}

type Tracker struct {
	mu           sync.Mutex
	started      int
	finished     int
	lastStarted  int
	lastFinished int
	entries      []GoroutineInfo
	wg           sync.WaitGroup
}

func (t *Tracker) Launch(label string, fn func()) {
	start := time.Now()

	t.mu.Lock()
	t.started++
	t.wg.Add(1)
	logWithTimestamp(StartColor("Starting goroutine"), label)
	t.mu.Unlock()

	go func() {
		defer func() {
			duration := time.Since(start)

			t.mu.Lock()
			t.finished++
			t.entries = append(t.entries, GoroutineInfo{
				Label:    label,
				Started:  start,
				Duration: duration,
			})
			logWithTimestamp(FinishColor("Finished goroutine"), label)
			t.mu.Unlock()

			t.wg.Done()
		}()

		fn()
	}()
}

func (t *Tracker) Summary() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.started == t.lastStarted && t.finished == t.lastFinished {
		return
	}

	t.lastStarted = t.started
	t.lastFinished = t.finished

	now := time.Now().Format("15:04:05")

	fmt.Println()
	fmt.Println(SummaryColor("--- Goroutine Summary ---"), "@", now)
	fmt.Println("Total started: ", t.started)
	fmt.Println("Total finished:", t.finished)
	fmt.Println("Currently running:", runtime.NumGoroutine())
	fmt.Println(SummaryColor("----------------------------"))
	fmt.Println()
}

func (t *Tracker) FinalReport() {
	t.mu.Lock()
	defer t.mu.Unlock()

	fmt.Println()
	fmt.Println(SummaryColor("✅ All Goroutines finished"))
	fmt.Println()

	fmt.Println(SummaryColor("=== Final Goroutine Report ==="))
	for _, entry := range t.entries {
		fmt.Printf("%-20s ran for %v\n", entry.Label, entry.Duration)
	}
	fmt.Println(SummaryColor("==============================="))
	fmt.Println()
}

func logWithTimestamp(prefix string, label string) {
	fmt.Printf("[%s] %s: %s\n",
		time.Now().Format("15:04:05.000"),
		prefix,
		label,
	)
}

func (t *Tracker) Wait() {
	t.wg.Wait()
}
