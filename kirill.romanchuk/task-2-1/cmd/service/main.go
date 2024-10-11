package main

import (
	"fmt"
	"math"
	"os"
)

const (
	LowerBound = 15
	UpperBound = 30
)

func readIntNum(message string, min int, max int) int {
	var num float64
	var resultNum int
	if message != "" {
		fmt.Print(message)
	}
	_, err := fmt.Scan(&num)
	if err != nil {
		panic("ошибка: некорректное значение")
	}
	resultNum = int(num)
	if num != float64(resultNum) {
		panic("ошибка: число должно быть целым")
	}
	if resultNum < min || resultNum > max {
		panic(fmt.Sprintf("ошибка: значение должно быть в диапазоне от %d до %d", min, max))
	}
	return resultNum
}

func readConditionAndTemperature() (string, int) {
	var condition string
	fmt.Print("Введите условие (>= или <=) и затем температуру (15-30): ")
	_, err := fmt.Scan(&condition)
	if err != nil {
		panic("ошибка: некорректное значение")
	}
	if condition != "<=" && condition != ">=" { //map?
		panic("Неверное условие: должно быть '<=' или '>='")
	}
	return condition, readIntNum("", LowerBound, UpperBound)
}

type Department struct {
	lowerBound int
	upperBound int
}

func (d *Department) manageTemperature() {
	k := readIntNum("Введите количество сотрудников (1-2000): ", 1, 2000)
	for j := 0; j < k; j++ {
		condition, temperature := readConditionAndTemperature()

		switch condition {
		case ">=":
			d.lowerBound = int(math.Max(float64(d.lowerBound), float64(temperature)))
		case "<=":
			d.upperBound = int(math.Min(float64(d.upperBound), float64(temperature)))
		}

		if d.lowerBound > d.upperBound {
			fmt.Println("Невозможно подобрать температуру для этого отдела\n", -1)
			break
		} else {
			fmt.Printf("Подходящая температура для отдела: %d\n", d.lowerBound)
		}
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()
	n := readIntNum("Введите количество отделов (1-2000): ", 1, 2000)
	departments := make([]Department, n)
	for i := 0; i < n; i++ {
		departments[i] = Department{lowerBound: LowerBound, upperBound: UpperBound}
	}
	for i := 0; i < n; i++ {
		departments[i].manageTemperature()
	}
}
