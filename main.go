package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, upper string) {
	length = len(name)
	upper = strings.ToUpper(name)
	return
} //naked return
func lenAndUpper2(name string) (int, string) {
	defer fmt.Println("I'm done")
	return len(name), strings.ToUpper(name)
} // nomal return

func repeatMe(words ...string) {
	defer fmt.Println("I'm done")
	fmt.Println(words)
}
func nomalFor(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

func forRange(numbers ...int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)

}

func main() {
	forRange(1, 2, 3, 4, 5, 6)

}
