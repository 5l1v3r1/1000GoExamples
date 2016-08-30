// Решить линейное уравнение A·x + B = 0, заданное своими коэффици-
// ентами A и B (коэффициент A не равен 0)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y float64
	fmt.Println("A·x + B = 0")
	x = util.NotNullNumber("Введите A")
	y = util.Number("Введите B")
	z := (0 - y) / x
	fmt.Println("x = ", z)
}
