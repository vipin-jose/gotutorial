package main

// Go routines are lightweight threads of execution that run concurrently with the main program.
// They are used to run multiple tasks simultaneously.
// concurrency is the ability to run multiple tasks simultaneously.
// it is not the same as parallelism.
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitgroup is used to wait for a collection of goroutines to finish executing.
var wg = sync.WaitGroup{}

// Mutex is a mutual exclusion lock. It is used to lock a memory location so that only one goroutine can access it at a time.
// var m = sync.Mutex{}
var m = sync.RWMutex{}

var dbData = []string{"data1", "data2", "data3", "data4", "data5"}
var result = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < 5; i++ {
		// sequential execution (Takes more than 5 seconds)
		// dbAccess(i)
		// concurrent execution (Takes less than 2 seconds)
		wg.Add(1) // Add 1 to the waitgroup
		go dbAccess(i)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Time taken for all goroutines to finish: ", time.Since(t0))
	fmt.Println("Result: ", result)
	channelTest()
}

func dbAccess(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Data from db: ", dbData[i])

	//result = append(result, dbData[i])
	// This can cause issues due to multiple goroutines accessing the same memory location
	// So we have to use a mutex to lock the memory location

	// Lock the memory location
	m.Lock()
	result = append(result, dbData[i])
	// Unlock the memory location
	m.Unlock()
	log()
	wg.Done() // Subtract 1 from the waitgroup
}

func log() {
	m.RLock()
	fmt.Printf("Data: %v\n", result)
	m.RUnlock()
}

//Channels are used to communicate between goroutines.
// Channels are used to send and receive data between goroutines.
// Channels are used to synchronize goroutines.
//Channels hold data
//Channels are thread-safe
// we can listen for data on a channel using the <- operator.

func channelTest() {
	var c = make(chan int)
	//c <- 1      // Send data to the channel
	// when we send data to a channel, the program will wait until the data is received by another goroutine.
	// If there is no goroutine to receive the data, the program will deadlock.
	//var i = <-c // Receive data from the channel and store it in i.
	//Channel is empty after this operation
	//fmt.Println(i)
	go process(c) // call process function as a goroutine
	//	fmt.Println(<-c) // value set by go routine. <-c waits for value to be set in goroutine
	for i := range c {
		fmt.Println(i)
	}

	var bufferedChannel = make(chan int, 5) // buffered channel with a capacity of 3
	go processBuffered(bufferedChannel)
	for i := range bufferedChannel {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
	//close(c)
}

func processBuffered(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting ProcessBuffered")
	//close(c)
}
