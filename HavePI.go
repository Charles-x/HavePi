package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random() float64 {
	//rand.Seed(time.Now().UnixNano())
	rdm := rand.Float64()
	return rdm
}

func main() {

	var N int = 2
	for {
		var cnt int = 0
		for i := 1; i <= N; i++ {
			x := random()
			y := random()


			if (x*x + y*y) <= 1 {
				cnt += 1
			}
		}
		vPi := 4.0 * float64(cnt) / float64(N)
		fmt.Println("Pi = ",vPi)
		N *= 2
		time.Sleep(time.Second)
	}
}
