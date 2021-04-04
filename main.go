package main

import (
	"fmt"
	"time"
)

func greet() {
	fmt.Printf("Hello World! %s", time.Now())
}

func main() {
	greet()
}
