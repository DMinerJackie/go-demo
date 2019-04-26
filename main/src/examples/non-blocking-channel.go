package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	/**
	A non-blocking send works similarly.
	Here msg cannot be sent to the messages channel,
	because the channel has no buffer and there is no receiver.
	Therefore the default case is selected.
	msg值无法发送到通道messages，因为一是它不是带缓冲的通道，二是没有任何的接收者
	*/
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signals", sig)
	default:
		fmt.Println("no activity")
	}
}
