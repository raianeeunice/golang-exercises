package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	host = make(chan bool, 2)
)
//ChopS struct
type ChopS struct{ 
	sync.Mutex 
}
//Philo struct
type Philo struct {
	number      int
	leftCS, rightCS  *ChopS
}

func (p Philo) eat() {
	for i := 0; i < 3; i++ {

		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.number+1)
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("finishing eating %d\n", p.number+1)

		p.leftCS.Unlock()
		p.rightCS.Unlock()

		<-host

	}
	wg.Done()
}

func main() {

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat()
	}

	wg.Wait()

}