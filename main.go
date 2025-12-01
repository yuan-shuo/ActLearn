package main

import (
	"fmt"
	"serve/utils"
	"time"
)

func main() {
	go func() {
		fmt.Println("a async print!")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("main process here~")

	fmt.Printf("1 + 2 = %d", utils.Add(1, 2))
}
