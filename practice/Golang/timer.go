package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	time.AfterFunc(time.Second*3, func() { fmt.Println("Hello") })

	time.Sleep(5 * time.Second)
}
