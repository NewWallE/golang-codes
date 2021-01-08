package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const numberOfGoRoutines = 200

//Create 200 go routines and demonstrate a race condition.
// Also show case the same using go run --race command in the terminal.
func main() {
	//raceCondition()
	raceConditionFixedWithMutex()

}

func raceConditionFixedWithMutex() {
	//Print some system properties at the start.
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	fmt.Println("[Start]Go Routines Count: ", runtime.NumGoroutine())
	counter := 0

	var mu sync.Mutex
	var ws sync.WaitGroup
	ws.Add(numberOfGoRoutines)

	for i := 0; i < numberOfGoRoutines; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			time.Sleep(time.Millisecond)
			counter = v
			mu.Unlock()
			ws.Done()
		}()
		fmt.Println("Go Routines Count: ", runtime.NumGoroutine())
		fmt.Println("running counter value: ", counter)
	}
	ws.Wait()
	fmt.Println("[End]Go Routines Count: ", runtime.NumGoroutine())
	fmt.Println("Counter Value: ", counter)

}
func raceCondition() {
	//Print some system properties at the start.
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	fmt.Println("[Start]Go Routines Count: ", runtime.NumGoroutine())
	counter := 0

	for i := 0; i < numberOfGoRoutines; i++ {
		go func() {
			// This counter++ is also running very fast and creating issue when run multiple times.
			// Since this is so fast that it is coming close to the value of numberOfGoRoutines.
			// Max goRoutine count is reaching upto 3. To increase the probability of dirty reads we need to
			// expand the below statement.
			//counter++
			v := counter
			v++
			//runtime.Gosched()
			// Putting time.sleep and sleeping for second 2 seconds. is drastically increasing the probability
			//of having race condition. The final counter value was coming out to be 0 or so.
			//time.Sleep(time.Second*2)
			time.Sleep(time.Millisecond)
			counter = v
		}()
		fmt.Println("Go Routines Count: ", runtime.NumGoroutine())
		fmt.Println("running counter value: ", counter)
	}

	fmt.Println("[End]Go Routines Count: ", runtime.NumGoroutine())
	fmt.Println("Counter Value: ", counter)

}
