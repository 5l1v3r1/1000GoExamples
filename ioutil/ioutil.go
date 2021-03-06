// Package ioutil утилиты для решения задач
package ioutil

import (
	"fmt"
	"math"
	"math/rand"
)

// Number ввод числа в stdin
func Number(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	return num
}

// Integer ввод целого числа в stdin
func Integer(msg string) int {
	fmt.Print(msg + " > ")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	return num
}

// UInteger ввод целого неотрицательного числа в stdin
func UInteger(msg string) uint {
	fmt.Print(msg + " > ")
	var num uint
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	return num
}

// NoNNumber ввод неотрицательного числа в stdin
func NoNNumber(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	for num < 0 {
		fmt.Println("Ввод неверных данных, введите неотрицательное число")
		_, err := fmt.Scanf("%f", &num)
		if err != nil {
			panic("Ввод неверных данных")
		}
	}
	return num
}

// NotNullNumber ввод ненулевого числа в stdin
func NotNullNumber(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	for num == 0 {
		fmt.Println("Ввод неверных данных, введите ненулевое число")
		_, err := fmt.Scanf("%f", &num)
		if err != nil {
			panic("Ввод неверных данных")
		}
	}
	return num
}

// GrNumber ввод градусов (0 < α < 360) в stdin
func GrNumber(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	for num < 0 || num > 360 {
		fmt.Println("Ввод неверных данных, введите 0 < α < 360")
		_, err := fmt.Scanf("%f", &num)
		if err != nil {
			panic("Ввод неверных данных")
		}
	}
	return num
}

// RadNumber ввод радиан (0 < α < 2·π) в stdin
func RadNumber(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	for num < 0 || num > 2*math.Pi {
		fmt.Println("Ввод неверных данных, введите 0 < α < 2·π")
		_, err := fmt.Scanf("%f", &num)
		if err != nil {
			panic("Ввод неверных данных")
		}
	}
	return num
}

// ModNumber приведение к модулю числа
func ModNumber(num float64) float64 {
	if num < 0 {
		num = -num
	}
	return num
}

// Symbol ввод string в stdin
func Symbol(msg string) string {
	fmt.Print(msg + " > ")
	var str string
	_, err := fmt.Scanf("%s", &str)
	if err != nil {
		panic("Ввод неверных данных")
	}
	return str
}

// RandomInt create random array []int
func RandomInt(n int) []int {
	list := rand.Perm(n)
	for i := 0; i < len(list); i++ {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
	return list
}

// RandomFlt create random array []float64
func RandomFlt(n int) []float64 {
	list := rand.Perm(n)
	var listEnd []float64
	for i := 0; i < len(list); i++ {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
	for k := 0; k < len(list); k++ {
		listEnd = append(listEnd, float64(list[k]))
	}
	return listEnd
}
