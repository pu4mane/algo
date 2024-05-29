package fizzbuzz

import "fmt"

func FizzBuzz() {
	fizz := "Fizz"
	buzz := "Buzz"
	fizzBuzz := fizz + buzz

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(fizzBuzz)
		} else if i%3 == 0 {
			fmt.Println(fizz)
		} else if i%5 == 0 {
			fmt.Println(buzz)
		} else {
			fmt.Println(i)
		}
	}
}
