package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var m sync.Mutex

func DivideByX(value chan string, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	for i := 1; i < 10; i++ {
		value <- name + strconv.Itoa(i/2)
		time.Sleep(time.Microsecond * time.Duration(3000))

	}
	m.Unlock()

}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	value1 := make(chan string, 10)
	value2 := make(chan string, 9)

	go DivideByX(value1, "Channel 1:", &wg)
	go DivideByX(value2, "Channel 2:", &wg)
	for i := 1; i < 10; i++ {
		valueX := <-value1
		fmt.Println(valueX)
	}
	// close(value1)
	for i := 1; i < 10; i++ {
		valueY := <-value2
		fmt.Println(valueY)
	}
	// close(value2)
	fmt.Println("waiting for go routines...")
	wg.Wait()
	fmt.Println("Done!")
}
