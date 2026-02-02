package main

import (
	"fmt"
	"sync"
)

func InvokeProcessor(wg *sync.WaitGroup, numProc int) {
	defer wg.Done()

	var procWg *sync.WaitGroup = new(sync.WaitGroup)
	procWg.Add(numProc)
	for procThread := range numProc {
		go runProcThread(procWg, procThread)
	}
	procWg.Wait()
	close(outputChan)
}

func runProcThread(procWg *sync.WaitGroup, procThread int) {
	defer procWg.Done()

	for {
		if currIPChar, ok := <-inputChan; ok {
			currInputChar := currIPChar.(byte)
			fmt.Printf("Processor Thread: %v reading character: %v\n", procThread, currInputChar)
			currOutputChar := currInputChar - 32
			fmt.Printf("Processor Thread: %v writing character: %v\n", procThread, currOutputChar)
			outputChan <- currOutputChar
		} else {
			continue
		}
	}
}
