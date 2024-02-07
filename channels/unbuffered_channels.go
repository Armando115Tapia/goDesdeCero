package channels

import (
	"fmt"
	"time"
)

func UnbufferedChannels() {
	ch := make(chan int) // Creating an unbuffered channel

	go func() {
		time.Sleep(time.Second) // Simulating some work
		ch <- 5                 // Sending a value to the chanel
	}()

	x := <-ch
	fmt.Println(x)

}
