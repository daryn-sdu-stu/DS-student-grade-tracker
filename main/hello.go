package main

import (
	"sync"
)

type Car struct {
	mu   sync.Mutex
	cars map[string]int
}

func (c *Car) inc(carName string) {
	c.mu.Lock()
	c.cars[carName]++
	defer c.mu.Unlock()
}

func incrementNumbers(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

/*func main() {
	log.Println("start main func")
	arrayOfNumber := rand.Perm(100)

	fmt.Println("arr from rnd numbers:  ", arrayOfNumber)
	fmt.Println("arr len:  ", len(arrayOfNumber))

	c := make(chan int)

	for i := 0; i < len(arrayOfNumber); i += 10 {
		go incrementNumbers(arrayOfNumber[i:i+10], c)
	}
	result := make([]int, 10)
	for i := 0; i < 10; i++ {
		result[i] = <-c
	}

	close(c)
	fmt.Println("result", result)

	//c := Car{
	//	cars: map[string]int{"audi": 0, "bmw": 0},
	//}
	//
	//var wg sync.WaitGroup
	//
	//doIncrement := func(carName string, n int) {
	//	for i := 0; i < n; i++ {
	//		c.inc(carName)
	//	}
	//
	//	wg.Done()
	//}
	//
	//wg.Add(3)
	//go doIncrement("audi", 3000)
	//go doIncrement("bmw", 3000)
	//go doIncrement("audi", 3000)
	//
	//wg.Wait()
	//fmt.Println(c.cars)
}*/
