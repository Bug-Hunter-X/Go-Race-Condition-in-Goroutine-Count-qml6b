# Go Race Condition in Goroutine Count

This repository demonstrates a common race condition in Go programs that use goroutines and `sync.WaitGroup` for counting.  The program attempts to increment a counter from multiple goroutines without proper synchronization, leading to an unreliable final count.

## Bug Description

The code uses a simple integer variable to track the number of completed goroutines.  Without any synchronization mechanisms (like mutexes or atomic operations), multiple goroutines accessing and modifying this variable concurrently cause data races, leading to unexpected and incorrect final counts.

## Solution

The solution uses an atomic counter to prevent race conditions.  This ensures that each increment operation happens atomically, preventing data races and giving a correct final count.  See `bugSolution.go` for the corrected code.
