package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Calculator provides basic math operations.
type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func (c *Calculator) Divide(a, b int) int {
	// BUG: no division-by-zero check
	return a / b
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	aStr := query.Get("a")
	bStr := query.Get("b")

	// BUG: no error handling for strconv.Atoi
	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)

	calc := &Calculator{}
	result := calc.Add(a, b)

	fmt.Fprintf(w, "Result: %d\n", result)
}

func main() {
	http.HandleFunc("/add", handler)
	fmt.Println("Server running on :8080")
	// BUG: no graceful shutdown, no timeout
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
