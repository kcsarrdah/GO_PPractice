package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello World %s! (iteration %d)\n", name, i+1)
		time.Sleep(300 * time.Millisecond)

	}
}

func main() {
	fmt.Println("Starting Routines ...")
	go sayHello("GoRoutine 1")
	go sayHello("Goroutine 2")

	sayHello("Main Thread") //not a thread(goroutine) so the main func will wait for it to finish

	time.Sleep(4 * time.Second)
	fmt.Println("Exiting Main Thread")

}
