package channels

import "fmt"

func Channels() {

	message := make(chan string)

	go func() { message <- "its message" }() // assigning values to channel

	msg := <-message // getting values from channels

	fmt.Println(msg)
}
