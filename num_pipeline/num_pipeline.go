package numpipeline

import "fmt"

func NumPipeline(num, sq chan int) {
	go func() {
		for x := 0; x <= 10; x++ {
			num <- x
		}
		close(num)
	}()

	go func() {
		for x := range num {
			sq <- x * x
		}
		close(sq)
	}()

	for x := range sq {
		fmt.Println(x)
	}
}
