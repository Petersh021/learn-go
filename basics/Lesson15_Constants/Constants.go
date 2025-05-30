package main

import "fmt"

const pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", pi, "Day")
	fmt.Println("Now you have", pi, "problems.")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
