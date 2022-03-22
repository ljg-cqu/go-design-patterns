package concurrency

import (
	"fmt"
	"time"
)

func Pipeline() {
	printA, printB, printC := make(chan bool), make(chan bool), make(chan bool)

	go func() {
		for range printA {
			fmt.Println("AAA")
			printB <- true
		}
	}()

	go func() {
		for range printB {
			fmt.Println("BBB")
			printC <- true
		}
	}()

	go func() {
		for range printC {
			fmt.Println("CCC")
		}
	}()

	t := time.NewTicker(time.Second)
	for i := 0; i < 10; i++ {
		<-t.C
		printA <- true
	}

	t.Stop()
}
