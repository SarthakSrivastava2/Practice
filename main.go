package main

import (
	"fmt"
	"sync"
)

var inputChan chan interface{}
var outputChan chan interface{}

var inputString, outputString string

var outputStrLock sync.RWMutex

func makeChan() {
	inputChan = make(chan interface{})
	outputChan = make(chan interface{})
}
func main() {
	var (
		numProc,
		numProd,
		numCon int
	)
	fmt.Printf("Enter number of producers\n")
	fmt.Scanf("%v", &numProd)
	fmt.Printf("Enter number of consumers\n")
	fmt.Scanf("%v", &numCon)
	fmt.Printf("Enter number of processors\n")
	fmt.Scanf("%v", &numProc)

	fmt.Printf("Enter input string\n")
	fmt.Scanf("%v", &inputString)

	var wg *sync.WaitGroup = new(sync.WaitGroup)
	makeChan()
	wg.Add(3)
	go InvokeProducer(wg, numProd)
	go InvokeProcessor(wg, numProc)
	go InvokeConsumer(wg, numCon)
	wg.Wait()
	fmt.Print()
}
