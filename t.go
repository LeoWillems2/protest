package main

//v1.1.9

import "fmt"

func Prt(s string) string {
	fmt.Println(s)
	return s + "!"
}

func main() {
	Prt("Sun")
}
