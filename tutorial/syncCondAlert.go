package main

import (
	"fmt"
	"sync"
	"time"
)

type Bank struct {
	money int
	cond  *sync.Cond
}

func (bank *Bank) Alert(value int) {
	bank.cond.L.Lock()
	for bank.money != value {
		bank.cond.Wait()
	}
	fmt.Printf("ALERT: Money is %v\n", bank.money)
	bank.cond.L.Unlock()
}

func main() {
	bank := &Bank{cond: sync.NewCond(&sync.Mutex{})}

	go bank.Alert(2)
	go bank.Alert(4)
	go bank.Alert(6) // On of these will never get cleaned,
	go bank.Alert(6) // it will wait for condition forever
	for {
		bank.cond.L.Lock()
		bank.money++
		fmt.Printf("Money is %v\n", bank.money)
		if bank.money == 6 {
			bank.cond.Signal() // Resume one waiter
		} else {
			bank.cond.Broadcast() // Resume all waiters
		}
		bank.cond.L.Unlock()
		time.Sleep(time.Second)
	}

}
