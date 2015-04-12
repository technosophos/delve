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

func testnext() {
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
}

func main() {
	d := make(chan interface{})
	testnext()
	go testgoroutine(d)
	<-d
	<-d
	fmt.Println("DONE")
}

func testgoroutine(d chan interface{}) {
	d <- 0
	close(d)
}
