package main

import "fmt"

type CircularBuffer[T any] struct {
	inputStream  <-chan T
	outputStream chan T
}

func NewCircularBuffer[T any](inputStream <-chan T, outputStream chan T) *CircularBuffer[T] {
	return &CircularBuffer[T]{
		inputStream:  inputStream,
		outputStream: outputStream,
	}
}

func (cb *CircularBuffer[T]) Run() {
	for v := range cb.inputStream {
		select {
		case cb.outputStream <- v:
		default:
			fmt.Printf("The buffer is full. Dropping the oldest value: %v\n", <-cb.outputStream)
			cb.outputStream <- v
		}
	}
	fmt.Println("The input stream was closed. Closing the output stream.")
	close(cb.outputStream)
}
