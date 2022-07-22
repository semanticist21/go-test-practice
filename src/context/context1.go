package context1

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	context, cancel := context.WithCancel(context.Background())
	go PrintEverySecond(context)

	time.Sleep(time.Second * 5)

	cancel()

	wg.Wait()
}

func PrintEverySecond(context context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-context.Done():
			wg.Done()
		case <-tick:
			fmt.Println("Ticked")
		}
	}
}
