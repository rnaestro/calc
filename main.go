package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateRomanNumber(s string) bool {
	romanNumbers := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	for _, n := range romanNumbers {
		if n == s {
			return true
		}
	}

	return false
}

func validateArabicNumber(n int) bool {
	return n > 0 && n < 11
}

func convertRomanToArabic(s string) int {
	return map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}[s]
}

func convertArabicToRoman(n int) string {
	if n/100 == 1 {
		return "C"
	}

	result := ""

	secondRanks := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}

	i := n / 10

	result += secondRanks[i]

	firstRanks := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	n = n % 10

	result += firstRanks[n]

	return result
}

func main() {
	var isValidRoman bool
	var firstOperand int
	var secondOperand int
	var result int

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите значение: ")

		inputString, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				break
			} else {
				fmt.Println("Ошибка: сбой при чтении операции")
				break
			}

		}

		// Конвертирование исходной строки в срез строк
		inputString = strings.TrimSpace(inputString)
		splitedInputString := strings.Split(inputString, " ")

		// Проверка количества вводимых аргументов
		if len(splitedInputString) < 3 {
			fmt.Println("Ошибка: строка не является математической операцией")
			break
		}
		if len(splitedInputString) > 3 {
			fmt.Println("Ошибка: строка должна представлять математическую операцию над двумя операнами")
			break
		}

		// Проверка является ли певый операнд римским числом входящим в множество [1, 10]
		isValidRoman = validateRomanNumber(splitedInputString[0])

		if isValidRoman {
			firstOperand = convertRomanToArabic(splitedInputString[0])
			if !validateRomanNumber(splitedInputString[2]) {
				fmt.Println("Ошибка: использование операндов разных систем счисления")
				break
			} else {
				secondOperand = convertRomanToArabic(splitedInputString[2])
			}
		} else {
			firstOperand, err = strconv.Atoi(splitedInputString[0])
			if err != nil {
				fmt.Printf("Ошибка: \"%s\" не является числом\n", splitedInputString[0])
				break
			}
			if !validateArabicNumber(firstOperand) {
				fmt.Printf("Ошибка: \"%s\" выходит за пределы допустимых значений. Операнд должен быть целым\nарабсим (I-X) или латинским(1-10) числом\n", splitedInputString[0])
				break
			}

			if validateRomanNumber(splitedInputString[2]) {
				fmt.Println("Ошибка: использование операндов разных систем счисления")
				break
			}

			secondOperand, err = strconv.Atoi(splitedInputString[2])
			if err != nil {
				fmt.Printf("Ошибка: \"%s\" не является числом\n", splitedInputString[2])
				break
			}
			if !validateArabicNumber(secondOperand) {
				fmt.Printf("Ошибка: \"%s\" выходит за пределы допустимых значений. Операнд должен быть целым \nарабсим (I-X) или латинским(1-10) числом\n", splitedInputString[2])
				break
			}
		}

		switch splitedInputString[1] {
		case "+":
			result = firstOperand + secondOperand
		case "-":
			result = firstOperand - secondOperand
		case "*":
			result = firstOperand * secondOperand
		case "/":
			result = firstOperand / secondOperand
		default:
			fmt.Printf("\"%s\" не является допустимым оператором\n", splitedInputString[1])
			break
		}

		if isValidRoman {
			if result < 1 {
				fmt.Println("Ошибка: результат операции с римскими числами не является положительным числом")
				break
			}
			fmt.Println(convertArabicToRoman(result))
		} else {
			fmt.Println(result)
		}
	}
}
