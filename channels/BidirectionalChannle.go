package channels

import "fmt"

func sendDataDos(ch chan<- int) {
	ch <- 42
}

func ChangeDireccionChannel() {
	ch := make(chan int)       // Declaring a bidirectional channel
	sendCh := (chan<- int)(ch) // Converting to send-only channel

	go sendDataDos(sendCh) // Sending data to the channel

	readCh := (<-chan int)(ch) // Converting to receive-only channel
	x := <-readCh              // Receiving data from the channel
	fmt.Println(x)

}
