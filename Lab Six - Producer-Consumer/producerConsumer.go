package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeBuffer is a thread-safe buffer for integers
type SafeBuffer struct {
	buffer []int      // Stores the numbers
	size   int        // Maximum number of items that can be stored
	mutex  sync.Mutex // Only one thread can access buffer at a time
	items  *sync.Cond // Wakes up consumers when items are available
	spaces *sync.Cond // Wakes up producers when space is available
}

// NewSafeBuffer creates a buffer with specified size
func NewSafeBuffer(size int) *SafeBuffer {
	sb := &SafeBuffer{
		buffer: make([]int, 0),
		size:   size,
	}
	// Both condition variables use the same mutex
	sb.items = sync.NewCond(&sb.mutex)
	sb.spaces = sync.NewCond(&sb.mutex)
	return sb
}

// Put adds an integer to the buffer
func (sb *SafeBuffer) Put(item int) {
	sb.mutex.Lock()
	// Wait if buffer is full
	for len(sb.buffer) >= sb.size {
		sb.spaces.Wait()
	}
	// Add item to end of buffer
	sb.buffer = append(sb.buffer, item)
	fmt.Printf("Produced %d, buffer size: %d/%d\n", item, len(sb.buffer), sb.size)
	sb.mutex.Unlock()
	// Tell consumers that an item is available
	sb.items.Signal()
}

// Get removes and returns an integer from the buffer
func (sb *SafeBuffer) Get() int {
	sb.mutex.Lock()
	// Wait if buffer is empty
	for len(sb.buffer) == 0 {
		sb.items.Wait()
	}
	// Take first item from buffer
	item := sb.buffer[0]
	sb.buffer = sb.buffer[1:]
	fmt.Printf("Consumed %d, buffer size: %d/%d\n", item, len(sb.buffer), sb.size)
	sb.mutex.Unlock()
	// Tell producers that space is available
	sb.spaces.Signal()
	return item
}

// Process simulates work on an integer value
func Process(value int) {
	fmt.Printf("Processing value %d\n", value)
	time.Sleep(time.Millisecond * 300)
}

// Producer creates integers and adds them to buffer
func producer(id int, buffer *SafeBuffer, wg *sync.WaitGroup, numItems int) {
	// Tell WaitGroup we're done when function ends
	defer wg.Done()
	for i := 0; i < numItems; i++ {
		// Create unique value based on producer ID
		value := id*100 + i
		buffer.Put(value)
		fmt.Printf("Producer %d created value %d\n", id, value)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Printf("Producer %d finished\n", id)
}

// Consumer takes integers from buffer and processes them
func consumer(id int, buffer *SafeBuffer, wg *sync.WaitGroup, numItems int) {
	// Tell WaitGroup we're done when function ends
	defer wg.Done()
	for i := 0; i < numItems; i++ {
		// Get value from buffer
		value := buffer.Get()
		// Process the value
		Process(value)
		fmt.Printf("Consumer %d processed value %d\n", id, value)
		time.Sleep(time.Millisecond * 150)
	}
	fmt.Printf("Consumer %d finished\n", id)
}

func main() {
	fmt.Println("Starting Producer-Consumer with Integers")

	// Configuration constants
	const (
		bufferSize   = 5 // Buffer can hold 5 items max
		numProducers = 2 // 2 producer threads
		numConsumers = 2 // 2 consumer threads
		itemsPer     = 6 // Each thread handles 6 items
	)

	var wg sync.WaitGroup
	buffer := NewSafeBuffer(bufferSize) // 5

	// Print configuration
	fmt.Printf("Buffer size: %d\n", bufferSize)
	fmt.Printf("Producers: %d, Consumers: %d\n", numProducers, numConsumers)
	fmt.Printf("Items per thread: %d\n\n", itemsPer)

	// Start producer threads
	wg.Add(numProducers)
	for i := 0; i < numProducers; i++ {
		go producer(i, buffer, &wg, itemsPer)
	}

	// Start consumer threads
	wg.Add(numConsumers)
	for i := 0; i < numConsumers; i++ {
		go consumer(i, buffer, &wg, itemsPer)
	}

	wg.Wait() // Wait for all threads to finish
	fmt.Println("All threads finished")
} // main
