package fibonacci

// func Fibonacci(n int) int {
// 	if n < 2 {
// 		return n
// 	}

// 	return Fibonacci(n-1) + Fibonacci(n-2)
// }

// var fiboCache = make(map[int]int)

// func Fibonacci(n int) int {
// 	if n < 2 {
// 		return n
// 	}

// 	if result, ok := fiboCache[n]; ok {
// 		return result
// 	}

// 	result := Fibonacci(n-1) + Fibonacci(n-2)

// 	fiboCache[n] = result

// 	return result
// }

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	a, b := 0, 1

	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}

	return b
}
