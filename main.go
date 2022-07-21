package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()

	fmt.Printf("Start Factory")
}

func MakeBody(tireCh chan *Car) {
	defer wg.Done()

	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports Car"
			tireCh <- car
		case <-after:
			close(tireCh)
			return
		}
	}
}

func InstallTire(tireCh chan *Car, paintCh chan *Car) {
	defer wg.Done()
	defer close(paintCh)

	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter Tire"
		paintCh <- car
	}
}

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
}
