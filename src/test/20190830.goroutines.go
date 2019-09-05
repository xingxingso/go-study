// Goroutines - Concurrency in Golang
// https://golangbot.com/goroutines/

// concurrency - No output from goroutine in Go - Stack Overflow
// https://stackoverflow.com/questions/28958192/no-output-from-goroutine-in-go

package main

import (
	"fmt"
	"time"
)

func main() {
	// test1_1();
	// test1_2();
	test1_3();
	// test2_1();
}

// Goroutines - Concurrency in Golang
// https://golangbot.com/goroutines/
func hello() {
	fmt.Println("Hello world goroutine")
}
func test1_1() {
	go hello()
	fmt.Println("main function")
	// go hello()
}
func test1_2() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
	// go hello()
}

func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func test1_3() {  
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond) // 1 a 2 3 b 4 c 5 d e main terminated 
    // time.Sleep(1000 * time.Millisecond) // 1 a 2 3 b main terminated
    fmt.Println("main terminated")
}

// concurrency - No output from goroutine in Go - Stack Overflow
// https://stackoverflow.com/questions/28958192/no-output-from-goroutine-in-go
func SayHello(done chan int) {
    for i := 0; i < 10; i++ {
        fmt.Print(i, " ")
    }
    if done != nil {
        done <- 0 // Signal that we're done
    }
    fmt.Print("\n")
}
func test2_1() {
	SayHello(nil) // Passing nil: we don't want notification here
	done := make(chan int)
	go SayHello(done)
	<-done // Wait until done signal arrives
}
