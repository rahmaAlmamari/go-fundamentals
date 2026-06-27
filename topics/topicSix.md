# Loops

In Go, there is only one looping construct: the for loop. However, it can be used in several different ways to achieve the behavior of while, do-while, and traditional loops from other languages.

| Loop Type               | Description                                               | Example                             |
| ----------------------- | --------------------------------------------------------- | ----------------------------------- |
| Traditional `for` loop  | Repeats a block of code a specific number of times        | `for i := 0; i < 5; i++`            |
| `for` as a `while` loop | Runs while a condition is true                            | `for x < 10`                        |
| Infinite loop           | Runs forever until stopped with `break` or `return`       | `for {}`                            |
| `for range` loop        | Iterates over arrays, slices, strings, maps, and channels | `for index, value := range numbers` |


## Traditional `for` Loop

```go 
for i := 0; i < 5; i++ {
	fmt.Println(i)
}

//output
// 0
// 1
// 2
// 3
// 4 

```

## `for` as a `while` Loop

```go 
count := 0

for count < 5 {
	fmt.Println(count)
	count++
}

//output
// 0
// 1
// 2
// 3
// 4

```

## Infinite Loop

```go 
for {
	fmt.Println("Running forever...")
}
```

This loop continues indefinitely until a break, return, or program termination occurs.

## `for range` Loop

1. Loop through a slice

```go 
numbers := []int{10, 20, 30}

for index, value := range numbers {
	fmt.Println(index, value)
}

//output
// 0 10
// 1 20
// 2 30

```

2. Ignore the index

```go
for _, value := range numbers {
	fmt.Println(value)
}

//output
// 10
// 20
// 30

```

3. Loop through a string

```go 
for index, char := range "Go" {
	fmt.Println(index, string(char))
}

//output
// 0 G
// 1 o

```

4. Loop through a map

```go
ages := map[string]int{
	"Ali": 25,
	"Sara": 22,
}

for name, age := range ages {
	fmt.Println(name, age)
}

//output
//Ali 25
//Sara 22

```

## Useful Keywords

| Keyword    | Description                                                         |
| ---------- | ------------------------------------------------------------------- |
| `break`    | Exits the loop immediately                                          |
| `continue` | Skips the current iteration and moves to the next one               |
| `range`    | Iterates over collections such as slices, arrays, strings, and maps |


1. Example using `continue`

```go
for i := 0; i < 5; i++ {
	if i == 2 {
		continue
	}

	fmt.Println(i)
}

//output
// 0
// 1
// 3
// 4

```

2. Example using `break`

```go
for i := 0; i < 5; i++ {
	if i == 3 {
		break
	}

	fmt.Println(i)
}

//output
// 0
// 1
// 2

```