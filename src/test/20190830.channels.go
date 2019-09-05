// Channels in Golang - Concurrency | golangbot.com
// https://golangbot.com/channels/

package main

import (
	"fmt"
	"time"
)

func main() {
	// test1_1()
	// test1_2()
	// test1_3()
	// test1_4()
	// test1_5()
	// test1_6()
	// test1_7()
	// test1_8()
	// test1_9()
	test1_10()
}

// Declaring channels
func test1_1() {
	var a chan int
	// a := make(chan int)  
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		// a := make(chan int)  
		fmt.Printf("Type of a is %T\n", a)
	}
}

func hello1_2() {
	fmt.Println("Hello world goroutinue")
}
func test1_2() {
	go hello1_2()
    time.Sleep(1 * time.Second)
    fmt.Println("main function")
}

func hello1_3(done chan bool) {  
    fmt.Println("Hello world goroutine")
    // time.Sleep(1 * time.Second)
    done <- true
}
func test1_3() {
	done := make(chan bool)
    go hello1_3(done)
    <-done
    fmt.Println("main function")
}

func hello1_4(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
    time.Sleep(4 * time.Second)
    fmt.Println("hello go routine awake and going to write to done")
    done <- true
}
func test1_4() {
	done := make(chan bool)
    fmt.Println("Main going to call hello go goroutine")
    go hello1_4(done)
    <-done
    fmt.Println("Main received data")
}

func calcSquares(number int, squareop chan int) {  
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit
        number /= 10
    }
    squareop <- sum
}
func calcCubes(number int, cubeop chan int) {  
    sum := 0 
    for number != 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
    }
    cubeop <- sum
} 
func test1_5() {  
    number := 589
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares(number, sqrch)
    go calcCubes(number, cubech)
    squares, cubes := <-sqrch, <-cubech
    fmt.Println("Final output", squares + cubes)
}

// Deadlock
func test1_6() {
	ch := make(chan int)
	ch <- 5
}

// Unidirectional channels
func sendData1_7(sendch chan<- int) {
    sendch <- 10
}
// func test1_7() {  
//     sendch := make(chan<- int)
//     go sendData1_7(sendch)
//     fmt.Println(<-sendch)
// }

func sendData1_8(sendch chan<- int) {  
	sendch <- 10
}
func test1_8() {  
    chnl := make(chan int)
    go sendData1_8(chnl)
    fmt.Println(<-chnl)
}

// Closing channels and for range loops on channels
func producer(chnl chan int) {  
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl)
}
func test1_9() {  
    ch := make(chan int)
    go producer(ch)
    for {
    	v, ok := <-ch
        if ok == false {
            break
        }
        fmt.Println("Received ", v, ok)
    }
}

func digits(number int, dchnl chan int) {  
    for number != 0 {
        digit := number % 10
        dchnl <- digit
        number /= 10
    }
    close(dchnl)
}
func calcSquares1_10(number int, squareop chan int) {  
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit
    }
    squareop <- sum
}
func calcCubes1_10(number int, cubeop chan int) {  
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit * digit
    }
    cubeop <- sum
}
func test1_10() {  
    number := 589
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares1_10(number, sqrch)
    go calcCubes1_10(number, cubech)
    squares, cubes := <-sqrch, <-cubech
    fmt.Println("Final output", squares+cubes)
}
