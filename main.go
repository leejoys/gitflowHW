package main

import (
	"fmt"
	"time"
)

func greet() {
	fmt.Printf("Hello Word! %s", time.Now())
}

func main() {
	greet()
}
