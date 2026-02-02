package main

import (
	"fmt"
	"sync"
	"time"
)

func InvokeProducer(wg *sync.WaitGroup, numProd int) {
	defer wg.Done()

	var prodWg *sync.WaitGroup = new(sync.WaitGroup)
	prodWg.Add(numProd)
	for prodThread := range numProd {
		go runProdThread(prodWg, prodThread)
	}
	prodWg.Wait()
	close(inputChan)
}

func runProdThread(prodWg *sync.WaitGroup, prodThread int) {

	defer prodWg.Done()
	outputStrLock.Lock()
	defer outputStrLock.Unlock()
	for len(inputString) > len(outputString) {
		currChar := len(outputString)
		if currChar < len(inputString) {
			prodChar := inputString[currChar]
			fmt.Printf("Producer Thread: %v reading character: %v\n", prodThread, prodChar)
			inputChan <- prodChar
			time.Sleep(100 * time.Millisecond)
		}
	}
}
