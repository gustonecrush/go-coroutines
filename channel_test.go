package golang_coroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Farhan Augustiansyah"
	}()

	data := <-channel
	fmt.Println(data)
}
