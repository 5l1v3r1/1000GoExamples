// Дни недели пронумерованы следующим образом: 1 — понедельник,
// 2 — вторник, ..., 6 — суббота, 7 — воскресенье. Дано целое число K , ле-
// жащее в диапазоне 1–365, и целое число N , лежащее в диапазоне 1–7. Оп-
// ределить номер дня недели для K -го дня года, если известно, что в этом
// году 1 января было днем недели с номером N
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y uint
	for x == 0 || x > 365 || x < 1 {
		x = util.UInteger("введите K, лежащее в диапазоне 1–365")
	}
	for y == 0 || y > 7 || y < 1 {
		y = util.UInteger("введите N, лежащее в диапазоне 1–7")
	}
	x = x % 7
	z := x + y
	if z > 7 {
		z = z - 7
	}
	fmt.Printf("номер K-го дня года: %v\n", z)
}
