package main

import "fmt"

func main() {
	//	fmt.Println(42)
	//	fmt.Printf("\n%b - %d\n", 42, 42)
	//	fmt.Printf("\noctal -\t %o hex %#x\n", 42, 42)
	for i := 60; i < 122; i++ {
		fmt.Printf("%o - %x - %d - %q\n", i, i, i, i)

	}
}
