package mypackage

import "fmt"

// Greet is an exported function
func Greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
