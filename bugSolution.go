```go
package main

import (

        "fmt"
        "sync"
        "sync/atomic"
)

func main() {
        var wg sync.WaitGroup
        var count int64
        const numRoutines = 1000

        for i := 0; i < numRoutines; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        atomic.AddInt64(&count, 1)
                }()
        }

        wg.Wait()
        fmt.Println("Final count:", count)
}
```