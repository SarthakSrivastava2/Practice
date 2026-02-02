package main

import (
	"fmt"
	"sync"
)

func InvokeConsumer(wg *sync.WaitGroup, numCon int) {
	defer wg.Done()

	var conWg *sync.WaitGroup = new(sync.WaitGroup)
	conWg.Add(numCon)
	for conThread := range numCon {
		go runConThread(conWg, conThread)
	}
	conWg.Wait()
}

func runConThread(conWg *sync.WaitGroup, conThread int) {
	defer conWg.Done()

	for {
		if currOPChar, ok := <-outputChan; ok {
			currOutputChar := currOPChar.(byte)
			fmt.Printf("Consumer Thread: %v reading character: %v\n", conThread, currOutputChar)
			outputString += string(currOutputChar)
			fmt.Printf("outputString: %v\n", outputString)
		} else {
			continue
		}
	}
}
