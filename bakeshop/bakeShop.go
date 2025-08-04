package main

import (
	"fmt"
	"sync"
	"time"
)

func preHeatOven() {
	fmt.Println("Starting to Preheat the Oven...")
	time.Sleep(3 * time.Second)
	fmt.Println("PreHeat finished!!")
}

func mixFlour() {
	fmt.Println("Mixing the Flour...")
	time.Sleep(3 * time.Second)
	fmt.Println("Flour mixed!!")
}

func makeChocoChips() {
	fmt.Println("Making the Choco Chips...")
	time.Sleep(3 * time.Second)
	fmt.Println("Choco Chips made!!")
}

func bake() {
	fmt.Println("Baking...")
	time.Sleep(3 * time.Second)
	fmt.Println("Baked!!")
}

func crushWalnuts() {
	fmt.Println("Crushing the Walnuts...")
	time.Sleep(3 * time.Second)
	fmt.Println("Walnuts crushed!!")
}

func bakeCookiesConcurrent() {
	fmt.Println("Starting to bake cookies concurrently...")
	startTime := time.Now()
	// WaitGroup to wait for all prep tasks to finish
	var wg sync.WaitGroup
	wg.Add(4) // We have 4 prep tasks

	// Start all prep tasks concurrently
	go func() {
		preHeatOven()
		wg.Done() // Signal this task is complete
	}()

	go func() {
		mixFlour()
		wg.Done()
	}()
	go func() {
		makeChocoChips()
		wg.Done()
	}()

	go func() {
		crushWalnuts()
		wg.Done()
	}()

	// Wait for ALL prep tasks to finish
	fmt.Println("⏳ Waiting for all prep tasks to complete...")
	wg.Wait()

	fmt.Println("🎯 All prep work done! Now we can bake!")
	bake()

	totalTime := time.Since(startTime)
	fmt.Printf("🕐 Concurrent baking took: %v\n", totalTime)

}

func bakeCookiesSequential() {
	fmt.Println("Starting to bake cookies sequentially...")
	startTime := time.Now()
	preHeatOven()
	mixFlour()
	makeChocoChips()
	crushWalnuts()
	bake()

	totalTime := time.Since(startTime)
	fmt.Printf("Baking cookies finished!! Time Taken: %v seconds\n", totalTime)
}

func main() {
	bakeCookiesConcurrent()
	bakeCookiesSequential()
	fmt.Println("\n🎉 Challenge complete! See the time difference?")

}
