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

func main() {
	totalLen, upper := lenAndUpper2("nwaaw")
	fmt.Println(totalLen, upper)
	repeatMe("ss", "ss2")

}
