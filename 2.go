package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(3 * time.Second)
}
