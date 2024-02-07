package channels

import (
	"fmt"
	"time"
)

func BufferedChannels() {
	ch := make(chan int, 2)

	go func() {
		time.Sleep(time.Second) // Simulating some work
		ch <- 5                 // Sending a value to the chanel
	}()
	ch <- 2

	x := <-ch
	fmt.Println(x)

	y := <-ch
	fmt.Println(y)

}
