package randnumsgenerator

import (
	"math/rand"
	"time"
)

func RandNumsGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(n)
		}
		close(out)
	}()
	return out
}
