// File: functions.go
// Purpose: A hands-on guide to functions in Go.
// Run: go run functions.go

package main

import (
	"errors"
	"fmt"
	"strings"
)

// Basic function: adds two integers.
func add(a int, b int) int {
	return a + b
}

// You can omit repeated type names for consecutive params of same type.
func subtract(a, b int) int {
	return a - b
}

// Multiple return values: common in Go.
func swap(a, b string) (string, string) {
	return b, a
}

// Named return values: the names become local variables and are returned when using a naked return.
// Use sparingly — explicit returns are clearer in many cases.
func divide(dividend, divisor float64) (quotient float64, err error) {
	if divisor == 0 {
		err = errors.New("divide by zero")
		return
	}
	quotient = dividend / divisor
	return // naked return uses named results
}

// Variadic function: accepts any number of ints.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Higher-order function: takes a function as an argument and applies it to each element.
func mapInts(xs []int, fn func(int) int) []int {
	out := make([]int, len(xs))
	for i, v := range xs {
		out[i] = fn(v)
	}
	return out
}

// Closure example: returns a function that captures 'start' from its environment.
func makeCounter(start int) func(int) int {
	count := start
	return func(delta int) int {
		count += delta
		return count
	}
}

// Recursion example: factorial (simple recursive implementation).
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Defer example: deferred calls execute when the function returns (LIFO order).
func deferExample(name string) {
	fmt.Println("entering:", name)
	defer fmt.Println("deferred 1 for", name)
	defer fmt.Println("deferred 2 for", name)
	fmt.Println("leaving:", name)
}

// Methods vs functions: define a type and attach methods.
type Person struct {
	FirstName string
	LastName  string
}

// Value receiver: does not modify the original.
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Pointer receiver: can modify the original.
func (p *Person) SetLastName(newLast string) {
	p.LastName = newLast
}

// Example of an anonymous function and immediate invocation.
var immediate = func(s string) string {
	return strings.ToUpper(s)
}("hello world")

func main() {
	fmt.Println("=== Basic math ===")
	fmt.Println("add(2,3) =", add(2, 3))
	fmt.Println("subtract(10,7) =", subtract(10, 7))

	fmt.Println("\n=== Swap (multiple returns) ===")
	a, b := swap("first", "second")
	fmt.Println("swapped:", a, b)

	fmt.Println("\n=== Divide with error handling ===")
	if q, err := divide(10, 2); err == nil {
		fmt.Println("10 / 2 =", q)
	} else {
		fmt.Println("error:", err)
	}
	if _, err := divide(1, 0); err != nil {
		fmt.Println("caught error:", err)
	}

	fmt.Println("\n=== Variadic sum ===")
	fmt.Println("sum(1,2,3,4) =", sum(1, 2, 3, 4))

	fmt.Println("\n=== mapInts (higher-order function) ===")
	nums := []int{1, 2, 3, 4}
	squared := mapInts(nums, func(x int) int { return x * x })
	fmt.Println("original:", nums)
	fmt.Println("squared:", squared)

	fmt.Println("\n=== Closures / stateful functions ===")
	counter := makeCounter(10)
	fmt.Println("counter +1 ->", counter(1))
	fmt.Println("counter +5 ->", counter(5))

	fmt.Println("\n=== Recursion (factorial) ===")
	fmt.Println("factorial(5) =", factorial(5))

	fmt.Println("\n=== Defer demonstration ===")
	deferExample("demo")

	fmt.Println("\n=== Methods on types ===")
	p := Person{"Ada", "Lovelace"}
	fmt.Println("FullName:", p.FullName())
	p.SetLastName("Byron")
	fmt.Println("After SetLastName:", p.FullName())

	fmt.Println("\n=== Anonymous immediate invocation ===")
	fmt.Println(immediate)

	fmt.Println("\n=== Function values and reassignment ===")
	var f func(int, int) int
	f = add
	fmt.Println("f(7,8) using assigned function:", f(7, 8))

	fmt.Println("\n=== Small exercises ===")
	fmt.Println("sum of nums slice via variadic:", sum(nums...))
	// Basic example of panic/recover (useful but handle with care)
	fmt.Println("\n=== Panic/Recover brief example ===")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered from panic:", r)
			}
		}()
		// Uncomment to see panic/recover in action:
		// panic("something bad happened")
		fmt.Println("panic not triggered in this run")
	}()

	// End of guide
	fmt.Println("\n--- End of function guide ---")
}