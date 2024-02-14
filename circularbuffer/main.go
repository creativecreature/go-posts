package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// producePrices is going to write 100 random stock prices to ch, and then close it.
func producePrices(ch chan<- int) {
	fmt.Println("The stock exchange has opened for the day.")
	for i := 0; i < 100; i++ {
		ch <- r.Intn(9999)
		// To simulare bursts with price changes we'll add a random delay between 0 and 4ms.
		time.Sleep(time.Duration(r.Intn(4)) * time.Millisecond)
	}
	fmt.Println("The stock exchange is closing for the day.")
	close(ch)
}

// calculateTrendLine is going to sleep for 2ms to simulate a calculation.
func calculateTrendLine() {
	// Sleep for 2ms to simulate that a calculation was performed.
	time.Sleep(2 * time.Millisecond)
}

func main() {
	originalStream := make(chan int)
	bufferedStream := make(chan int, 3)

	// Initialize the circular buffer with the input and output streams.
	cb := NewCircularBuffer(originalStream, bufferedStream)
	go cb.Run()

	// Simulate a stream of data that is going to produce stock prices at a high phase.
	go producePrices(originalStream)

	for v := range bufferedStream {
		calculateTrendLine()
		fmt.Printf("Updated the trend line with value: %v\n", v)
	}
}
