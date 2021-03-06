// Дано вещественное число X (| X | < 1) и целое число N (> 0). Найти значение выражения
// X – X^3 /3 + X^5 /5 – ... + (–1)N · X^2·N+1 /(2· N +1).
// Полученное число является приближенным значением функции arctg в
// точке X

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		x float64
		n int
	)

	x = 2
	for x >= 1 {
		x = ioutil.NotNullNumber("X")
	}

	ioutil.ModNumber(x)

	n = -1
	for n < 0 {
		n = ioutil.Integer("N")
	}
	fmt.Println("X – X^3 /3 + X^5 /5 – ... + (–1)N · X^2·N+1 /(2· N +1)")
	fmt.Printf("%v\n", foo1(n, x))
}

func foo1(n int, x float64) float64 {
	k := x
	a := 1.0
	for i := 2; i <= n; i++ {
		if i%2 != 0 {
			a = -a
			k = k + a*(math.Pow(x, float64(i))/float64(i))
		}
	}
	return k
}
