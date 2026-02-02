package main

import (
	"fmt"
	"log"
	"sync"
)

// Producer: generates a range of numbers and sends them to input channel
func producer(id, start, end int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		log.Printf("[producer-%d] -> %d", id, i)
		out <- i
	}
}

// Processor: reads from input, processes, sends to output (worker)
func processor(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		res := n * 2 // simple processing
		log.Printf("  [processor-%d] %d -> %d", id, n, res)
		out <- res
	}
}

// Consumer: reads final results (worker)
func consumer(id int, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		log.Printf("    [consumer-%d] Consumed: %d", id, n)
	}
}

func main() {
	// Tunables: adjust pool sizes here
	numProducers := 2
	numProcessors := 3
	numConsumers := 2

	// Channels (buffered to allow some pipelining/backpressure)
	input := make(chan int, 16)
	output := make(chan int, 16)

	var wgProd, wgProc, wgCons sync.WaitGroup

	// --- Start Producers ---
	// Example: split work ranges across producers
	wgProd.Add(numProducers)
	totalItems := 20
	chunk := totalItems / numProducers
	remainder := totalItems % numProducers
	start := 1
	for p := 1; p <= numProducers; p++ {
		size := chunk
		if p == numProducers {
			size += remainder // assign remainder to last producer
		}
		end := start + size - 1
		go producer(p, start, end, input, &wgProd)
		start = end + 1
	}

	// Close input when all producers are done
	go func() {
		wgProd.Wait()
		close(input)
	}()

	// --- Start Processors (worker pool) ---
	wgProc.Add(numProcessors)
	for i := 1; i <= numProcessors; i++ {
		go processor(i, input, output, &wgProc)
	}

	// Close output when all processors are done
	go func() {
		wgProc.Wait()
		close(output)
	}()

	// --- Start Consumers (worker pool) ---
	wgCons.Add(numConsumers)
	for i := 1; i <= numConsumers; i++ {
		go consumer(i, output, &wgCons)
	}

	// Wait for consumers to finish draining output
	wgCons.Wait()

	fmt.Println("Pipeline complete")
}
