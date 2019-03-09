package main

import "fmt"
import "time"

func main() {
	requests := make(chan int, 5)
	for i := 1; i<=5; i++ {
		
		requests <- i	
	}


	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("=========")
	
	
	burstyLimiter := make(chan time.Time, 3)


	//lian xu zhu ru 3 ge shuju
	for i:=0; i<3; i++ {
		burstyLimiter <- time.Now()	
	}

	go func(){
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 8)
	for i :=1; i<=8 ;i++{
		burstyRequests <- i	
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
