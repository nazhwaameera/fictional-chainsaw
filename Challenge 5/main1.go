package main

import (
	"fmt"
	"sync"
)

func main() {
	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go printInterface(i, data1, &wg)
	}

	for j := 1; j <= 4; j++ {
		wg.Add(1)
		go printInterface(j, data2, &wg)
	}

	wg.Wait()
}

func printInterface(index int, data []interface{}, wg *sync.WaitGroup) {
	fmt.Println(data, index)
	wg.Done()
}
