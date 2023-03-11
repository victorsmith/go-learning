package main

import (
	"fmt"
	"time"
)

func print10() {
	for i := 0; i < 5; i++ {
		// Printf has formatting directives
		// Println works more like js's console.log
		fmt.Println("Print 10:", i)
	}
}

func main() {
	go print10()
	time.Sleep(3 * time.Second)
}
