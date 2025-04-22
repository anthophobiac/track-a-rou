# track-a-rou
(..tine)

Tiny tool to track goroutines — logs when goroutines start and finish, shows live summaries, and prints a final  report once all goroutines are completed.

## Features

- Track how many goroutines are started and finished
- Show summary when activity changes
- Colour-coded terminal output using [`fatih/color`](https://github.com/fatih/color)
- Uses `sync.WaitGroup` to wait for all goroutines to complete

## Getting Started

```bash
git clone https://github.com/anthophobiac/track-a-rou.git
cd track-a-rou
go mod tidy
go run main.go
```

## Example Output

```bash

[22:33:20.123] Starting goroutine: worker-0
[22:33:21.456] Finished goroutine: worker-0

--- Goroutine Summary --- @ 22:33:21.456
Total started:  1
Total finished: 1
Currently running: 0
----------------------------

✅ All Goroutines finished

=== Final Goroutine Report ===
worker-0            ran for 1.333s
===============================
```