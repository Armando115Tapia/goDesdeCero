package channels

import (
	"fmt"
	"time"
)

func sendData(ch chan int) {
	ch <- 42
}

func readData(ch chan int) {
	x := <-ch
	fmt.Println(x)
}

func BidirectionalChannel() {
	ch := make(chan int, 1)

	go readData(ch)
	go sendData(ch)

	time.Sleep(time.Second)
}
