package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}
    
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("1st func")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg,mut)
	go func(wg *sync.WaitGroup,  m *sync.Mutex){
		fmt.Println("2nd func")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg,mut)
	go func(wg *sync.WaitGroup,  m *sync.Mutex){
		fmt.Println("3rd func")
		m.Lock()
		score = append(score, 3)  // Locking the memory to avoid race condition or competition for the memory
		m.Unlock()
		wg.Done()
	}(wg,mut)

	wg.Wait()

}