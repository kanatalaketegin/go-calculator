package main

import (
	"fmt"
	"strings"
)

func main() {
	var numberOne, numberTwo, operation string
	fmt.Scan(&numberOne, &operation, &numberTwo)
}

func calculate(num1, num2 int, operation string) int {

	switch operation {
	case "+":
		num1 = num1 + num2
		break
	case "-":
		num1 = num1 - num2
		break
	case "*":
		num1 = num1 * num2
		break
	case "/":
		if num1 != 0 || num2 != 0 {
			num1 = num1 / num2
		} else {
			fmt.Println("error division the 0")
		}
		break
	default:
		fmt.Println("error")
	}
	return num1
}

func intToRoman(num int) string {
	mapping := make(map[int]string)
	mapping[1] = "I"
	mapping[4] = "IV"
	mapping[5] = "V"
	mapping[9] = "IX"
	mapping[10] = "X"
	mapping[40] = "XL"
	mapping[50] = "L"
	mapping[90] = "XC"
	mapping[100] = "C"
	mapping[400] = "CD"
	mapping[500] = "D"
	mapping[900] = "CM"
	mapping[1000] = "M"

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var result string

	for _, value := range nums {
		if num >= value {
			r := num / value
			result = result + strings.Repeat(mapping[value], r)
			num = num % value
		}
	}
	return result
}

func romanToInt(num string) int {
	result := 0

	for item := 0; item < len(num); item++ {
		if num[item] == 'M' {
			result += 1000
		}
		if num[item] == 'D' {
			result += 500
		}
		if num[item] == 'C' {
			if item+1 < len(num) && num[item+1] == 'D' {
				result += 400
				item++
				continue
			}
			if item+1 < len(num) && num[item+1] == 'M' {
				result += 900
				item++
				continue
			}
			result += 100
		}
		if num[item] == 'L' {
			result += 50
		}
		if num[item] == 'X' {
			if item+1 < len(num) && num[item+1] == 'L' {
				result += 40
				item++
				continue
			}
			if item+1 < len(num) && num[item+1] == 'C' {
				result += 90
				item++
				continue
			}
			result += 10
		}
		if num[item] == 'V' {
			result += 5
		}
		if num[item] == 'I' {
			if item+1 < len(num) && num[item+1] == 'V' {
				result += 4
				item++
				continue
			}
			if item+1 < len(num) && num[item+1] == 'X' {
				result += 9
				item++
				continue
			}
			result += 1
		}
	}
	return result

}
