package main

import "fmt"
import "time"


func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now, secs, nanos)


	millis := nanos / 1000000
	fmt.Println(millis)


	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}


