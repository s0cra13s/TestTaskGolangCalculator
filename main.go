package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func errExit() {
	ErrNotValid := errors.New("Not valid")
	fmt.Println(ErrNotValid)
	os.Exit(1)
}

func isRoma(input string) bool {
	if string(input[0]) == "I" || string(input[0]) == "V" || string(input[0]) == "X" {
		return true
	}
	return false
}

func roma_to_int(input string) int {
	result := 0
	for i := 0; i < len(input); i++ {
		switch {
		case string(input[i]) == "X":
			result += 10
			break
		case string(input[i]) == "V":
			result += 5
			break
		case string(input[i]) == "I":
			result += 1
			break
		}
	}
	for i := 1; i < len(input); i++ {
		if (string(input[i]) == "V" || string(input[i]) == "X") && string(input[i-1]) == "I" {
			result -= 1 + 1
		}
	}
	return result
}

func int_to_roma(input int) string {
	var str_romans = []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var values = []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	var result string
	for i := 0; i < len(values); i++ {
		for input-values[i] >= 0 {
			result += str_romans[i]
			input -= values[i]
		}
	}
	return result
}

func main() {
	var a, b, _, result, operator_count int
	var as, bs, operator string
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	for i := 0; i < len(input); i++ {
		if (string(input[i]) == "*" || string(input[i]) == "/" || string(input[i]) == "+" || string(input[i]) == "-") && operator_count == 0 {
			operator = string(input[i])
			operator_count += 1
		} else if (string(input[i]) == "*" || string(input[i]) == "/" || string(input[i]) == "+" || string(input[i]) == "-") && operator_count != 0 {
			errExit()
		}
	}
	switch {
	case operator == "*":
		as = strings.Split(input, "*")[0]
		bs = strings.Split(input, "*")[1]
	case operator == "/":
		as = strings.Split(input, "/")[0]
		bs = strings.Split(input, "/")[1]
	case operator == "+":
		as = strings.Split(input, "+")[0]
		bs = strings.Split(input, "+")[1]
	case operator == "-":
		as = strings.Split(input, "-")[0]
		bs = strings.Split(input, "-")[1]
	default:
		errExit()
	}
	aRoma := isRoma(as)
	bRoma := isRoma(bs)
	if aRoma == true && bRoma == true {
		a = roma_to_int(as)
		b = roma_to_int(bs)
	} else if aRoma == false && bRoma == false {
		a, _ = strconv.Atoi(as)
		b, _ = strconv.Atoi(bs)
	} else {
		errExit()
	}
	if a > 10 || b > 10 || a < 1 || b < 1 {
		errExit()
	} else {
		switch {
		case operator == "*":
			result = a * b
		case operator == "/":
			result = a / b
		case operator == "+":
			result = a + b
		case operator == "-":
			result = a - b
		default:
			errExit()
		}
	}
	if aRoma == true && result > 0 {
		print(int_to_roma(result))
	} else if aRoma == true && result <= 0 {
		errExit()
	} else {
		print(result)
	}

}
