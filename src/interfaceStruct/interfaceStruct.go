package interfacestruct

import (
	"fmt"
)

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct {
}

func (alarm *Alarm) OnFire() {
	fmt.Println("Onfire")
}

func Go() {
	mail := &Mail{}
	var listener EventListener = &Alarm{}
	mail.Register(listener)
	mail.OnRecv()
}
