package main

import (
	"fmt"
	"sync"
)

func main() {
	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}
	n := 1
	count1 := 0
	count2 := 0

	var wg sync.WaitGroup
	var m sync.Mutex

	for n <= 8 {

		if n%2 == 0 {
			wg.Add(1)
			count1++
			Turn(&n, &m)
			go printInterface(count1, data1, &wg, &m)
			// fmt.Println(n)
		} else {
			wg.Add(1)
			count2++
			Turn(&n, &m)
			go printInterface(count2, data2, &wg, &m)
			// fmt.Println(n)
		}
	}
	wg.Wait()
}

func printInterface(index int, data []interface{}, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	fmt.Println(data, index)
	wg.Done()
	m.Unlock()
}

func Turn(n *int, m *sync.Mutex) {
	m.Lock()
	*n++
	m.Unlock()
}
