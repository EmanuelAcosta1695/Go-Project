package main

import "fmt"

func main() {
	server := NewAPIserver(":3000")
	server.Run()
	fmt.Println("Yeah!")
}
