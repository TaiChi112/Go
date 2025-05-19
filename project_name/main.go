package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(a int, b int) int {
	return a + b
}
func swap(a, b string) (string, string) {
	return b, a
}
func main() {
	var a, b string
	a, b = "Alice", "Bob"
	fmt.Printf("%s %s\n", a, b)
	a, b = swap("Alice", "Bob")
	fmt.Println(a, b)
	fmt.Println(add(1, 2))
	fmt.Println("Random number:", rand.Intn(100))
	fmt.Println("The time is", time.Now())
	fmt.Println(math.Pi)
}
