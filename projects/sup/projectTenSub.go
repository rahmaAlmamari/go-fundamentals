package sup

import "fmt"

// Exported function (public)
func Add(a int, b int) int {
	return a + b
}

// Unexported function (private to package)
func logResult(result int) {
	fmt.Println("Result is:", result)
}
