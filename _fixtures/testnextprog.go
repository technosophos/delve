// fix lines
package main

import (
	"fmt"
	"time"
)

func sleepytime() {
	time.Sleep(5 * time.Millisecond)
}

func helloworld() {
	fmt.Println("Hello, World!")
}

func testnext(d chan interface{}) {
	var (
		j = 1
		f = 2
	)

	for i := 0; i <= 5; i++ {
		j += j * (j ^ 3) / 100

		if i == f {
			fmt.Println("foo")
			break
		}

		sleepytime()
	}

	helloworld()
	close(d)
}

func main() {
	d := make(chan interface{})
	go testnext(d)
	<-d
	fmt.Println("DONE")
}
