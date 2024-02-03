# Stacking Defers in Go

In Go, deferred function calls are managed through a stack mechanism. When a function is executed, any deferred calls within that function are pushed onto a stack. As the function completes and returns, the deferred calls are executed in a last-in-first-out (LIFO) order.

## Example

```go
package main

import "fmt"

func main() {
    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}

```output
counting
done
9
8
7
6
5
4
3
2
1
0
