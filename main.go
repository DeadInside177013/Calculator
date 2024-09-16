package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romeToArab = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "L": 50, "C": 100,
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	line = strings.ReplaceAll(line, " ", "")

	op, err := symbol(line)
	if err != nil {
		fmt.Println(err)
		return
	}

	arguments := strings.Split(line, op)
	if len(arguments) != 2 {
		fmt.Println("incorrect format of expression!")
		return
	}

	argument1 := arguments[0]
	argument2 := arguments[1]

	var a, b int
	var isRoman bool

	if IsRome(argument1) && IsRome(argument2) {
		a = romeToArab[argument1]
		b = romeToArab[argument2]
		isRoman = true
	} else {
		a, err = strconv.Atoi(argument1)
		if err != nil {
			fmt.Println("conflicting arguments")
			return
		}
		b, err = strconv.Atoi(argument2)
		if err != nil {
			fmt.Println("conflicting arguments")
			return
		}
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		fmt.Println("Numbers cannot be less 1 or more than 10")
		return
	}

	result, err := numOperarion(a, b, op)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isRoman {
		if result <= 0 {
			fmt.Println("In roman numerals, the result cannot be less than or equal ro zero")
		} else {
			res := arabTorome(result)
			fmt.Println("Result: %s\n", res)
		}
	} else {
		fmt.Println("Result: %d\n", result)
	}

}

func symbol(line string) (string, error) {
	switch {
	case strings.Contains(line, "+"):
		return "+", nil
	case strings.Contains(line, "-"):
		return "-", nil
	case strings.Contains(line, "*"):
		return "*", nil
	case strings.Contains(line, "/"):
		return "/", nil
	default:
		return "", fmt.Errorf("can't find operator")
	}
}

func numOperarion(a, b int, op string) (num int, err error) {
	switch op {
	case "+":
		num = a + b
	case "-":
		num = a - b
	case "*":
		num = a * b
	case "/":
		num = a / b
	default:
		return 0, fmt.Errorf("unsupported operator")
	}
	return num, nil
}

func IsRome(line string) bool {
	if len(line) == 0 {
		return false
	}
	for _, i := range line {
		if _, ok := romeToArab[string(i)]; ok {
			return true
		}
	}
	return false
}

func arabTorome(num int) string {
	roman := ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}
	return roman
}
