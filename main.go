package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const op = "+-*/"

func main() {
	fmt.Print("Введите арифметическую операцию (например, '2 + 3'):\n")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Удаление символа новой строки

	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Ошибка! Формат арифметической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	operator := parts[1]
	if !strings.ContainsAny(operator, op) {
		panic("Ошибка! Второй элемент ввода не является оператором.")
	}

	fmt.Println(result(parts[0], operator, parts[2]))
}

func result(x, a, y string) string {
	parsSign(a)
	first, err1 := strconv.Atoi(x)
	second, err2 := strconv.Atoi(y)
	if err1 == nil && err2 == nil {
		return strconv.Itoa(parsNumber(first, second, a))
	} else {
		return parsRome(x, a, y)
	}
}

func parsSign(a string) string {
	if strings.ContainsAny(a, op) {
		return a
	}
	panic("Ошибка! Неверный арифметический знак (оператор). Допустимо: +-*/")
}

func parsNumber(x, y int, a string) int {
	var res int
	if x > 0 && x < 11 && y > 0 && y < 11 {
		switch a {
		case "+":
			res = x + y
		case "-":
			res = x - y
		case "*":
			res = x * y
		case "/":
			res = x / y
		}
		return res
	} else {
		panic("Ошибка! Ожидаются только арбские либо только рисмские числа от 1 до 10")
	}
}

func parsRome(x, a, y string) string {
	numRome := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	firstRome, ok1 := numRome[x]
	secondRome, ok2 := numRome[y]
	if ok1 && ok2 {
		num := parsNumber(firstRome, secondRome, a)
		if num < 0 {
			panic("Ошибка! В римской системе нет отрицательных чисел.")
		}
		symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		var result strings.Builder
		for i := 0; i < len(symbols); i++ {
			for num >= values[i] {
				result.WriteString(symbols[i])
				num -= values[i]
			}
		}
		return result.String()
	} else {
		panic("Ошибка! Ожидаются только арбские либо только рисмские числа которые не больше 10")
	}
}
