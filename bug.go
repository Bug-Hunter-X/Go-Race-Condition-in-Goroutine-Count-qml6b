```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        const numRoutines = 1000

        for i := 0; i < numRoutines; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        count++
                }()
        }

        wg.Wait()
        fmt.Println("Final count:", count)
}
```