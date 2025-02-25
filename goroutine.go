package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello() {
	fmt.Println("Hello, Goroutine!")
}

func worker(done chan bool) {
	fmt.Println("Working...")
	done <- true
}

func worker2(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mengurangi counter WaitGroup setelah selesai
	fmt.Printf("Worker %d is working\n", id)
}

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock() // Kunci agar hanya satu goroutine yang bisa mengakses variabel ini
	counter++
	mutex.Unlock() // Buka kunci setelah selesai
}

func main() {
	go sayHello()
	time.Sleep(time.Second)

	fmt.Println("=====================")
	done := make(chan bool)
	go worker(done)
	<-done // Menunggu sinyal dari goroutine

	fmt.Println("=====================")
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Tambah counter sebelum goroutine dijalankan
		go worker2(i, &wg)
	}
	wg.Wait() // Tunggu semua goroutine selesai
	fmt.Println("All workers done!")

	fmt.Println("=====================")
	var wg2 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go increment(&wg2)
	}
	wg2.Wait()
	fmt.Println("Counter:", counter) // Output harus 10 tanpa race condition
}
