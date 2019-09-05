package main

import (
	"fmt"
	"sync"
)

func main() {
	// test1_1()
	// test1_2()
	// test1_3()
	// test1_4()
	// test1_5()
	test1_6()
}

// do well
func test1_1() {
	go allocateJob(10)
	done := make(chan bool)
	go printResult(done)
	createWorkerPool(2)
	<-done
}

// just print 1 to 10
func test1_2() {
	go allocateJob(10)
	for v := range jobs {
		fmt.Println(v.id)
	}
	done := make(chan bool)
	go printResult(done)
	createWorkerPool(2)
	<-done
}

// error
func test1_3() {
	go allocateJob(10)
	createWorkerPool(2)
	done := make(chan bool)
	go printResult(done)
	<-done
}

// do well
func test1_4() {
	go allocateJob(10)
	go createWorkerPool(2)
	done := make(chan bool)
	go printResult(done)
	<-done
}

// do well
func test1_5() {
	go allocateJob(10)
	done := make(chan bool)
	go printResult(done)
	go createWorkerPool(2)
	<-done
}

// error
func test1_6() {
	go allocateJob(10)
	done := make(chan bool)
	go printResult(done)
	<-done
	createWorkerPool(2)
}

type Job struct {
	id int 
	number int
}
type Result struct {
	job Job
	res int
}

var jobs = make(chan Job)
var results = make(chan Result)

func allocateJob(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		jobs <- Job{i+1, (i+1)*(i+1)}
	}
	close(jobs)
}

func printResult(done chan bool) {
	for v := range results {
		fmt.Println(v.job.id, v.job.number, v.res)
	}

	done <- true
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup) {
	for v := range jobs {
		results <- Result{v, doWork(v.number)}
	}
	wg.Done()
}

func doWork(number int) int {
	return number * number
}
