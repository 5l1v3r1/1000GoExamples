// Дано значение температуры T в градусах Цельсия. Определить значе-
// ние этой же температуры в градусах Фаренгейта. Температура по Цельсию
// Tc и температура по Фаренгейту Tf связаны следующим соотношением:
// Tc = (Tf – 32)·5/9
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x float64
	x = util.Number("Введите температуру T в градусах Цельсия")
	fmt.Println("Температура по Фаренгейту:\t", (x-32)*5/9)
}
