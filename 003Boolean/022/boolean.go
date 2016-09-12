// Дано трехзначное число. Проверить истинность высказывания:
// «Цифры данного числа образуют возрастающую или убывающую после-
// довательность»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a int
	var x, y bool
	for a > 999 || a < 100 {
		a = ioutil.Integer("целое трехзначное число A")
	}
	x = (a/100) < ((a-(100*(a/100)))/10) && ((a-(100*(a/100)))/10) < (a%10)
	y = (a/100) > ((a-(100*(a/100)))/10) && ((a-(100*(a/100)))/10) > (a%10)
	fmt.Println(x || y)
}
