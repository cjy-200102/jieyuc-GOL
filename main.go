package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello("cjy")
	fmt.Println("Hello The Game of life")
	time.Sleep(1 * time.Second)
}

func sayHello(name string) {
	fmt.Println("hello", name)
}
