// Дано целое число A . Проверить истинность высказывания: «Число A
// является четным»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x int
	x = ioutil.Integer("целое число A")
	fmt.Println(x%2 == 0)
}
