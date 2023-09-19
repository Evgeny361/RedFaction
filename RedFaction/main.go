package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var rome = []string{
		"",
		"I",
		"II",
		"III",
		"IV",
		"V",
		"VI",
		"VII",
		"VIII",
		"IX",
		"X",
		"XI",
		"XII",
		"XIII",
		"XIV",
		"XV",
		"XVI",
		"XVII",
		"XVIII",
		"XIX",
		"XX",
		"XXI",
		"XXII",
		"XXIII",
		"XXIV",
		"XXV",
		"XXVI",
		"XXVII",
		"XXVIII",
		"XXIX",
		"XXX",
		"XXXI",
		"XXXII",
		"XXXIII",
		"XXXIV",
		"XXXV",
		"XXXVI",
		"XXXVII",
		"XXXVIII",
		"XXXIX",
		"XL",
		"XLI",
		"XLII",
		"XLIII",
		"XLIV",
		"XLV",
		"XLVI",
		"XLVII",
		"XLVIII",
		"XLIX",
		"L",
		"LI",
		"LII",
		"LIII",
		"LIV",
		"LV",
		"LVI",
		"LVII",
		"LVIII",
		"LIX",
		"LX",
		"LXI",
		"LXII",
		"LXIII",
		"LXIV",
		"LXV",
		"LXVI",
		"LXVII",
		"LXVIII",
		"LXIX",
		"LXX",
		"LXXI",
		"LXXII",
		"LXXIII",
		"LXXIV",
		"LXXV",
		"LXXVI",
		"LXXVII",
		"LXXVIII",
		"LXXIX",
		"LXXX",
		"LXXXI",
		"LXXXII",
		"LXXXIII",
		"LXXXIV",
		"LXXXV",
		"LXXXVI",
		"LXXXVII",
		"LXXXVIII",
		"LXXXIX",
		"XC",
		"XCI",
		"XCII",
		"XCIII",
		"XCIV",
		"XCV",
		"XCVI",
		"XCVII",
		"XCVIII",
		"XCIX",
		"C",
	}
	var a, b, c string
	var str string
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	strSlice := strings.Fields(str)
	if len(strSlice) > 3 {
		fmt.Println("Некорректные данные(слишком много символов)")
		return
	}
	a = strings.ToUpper(strSlice[0])
	b = strSlice[1]
	c = strings.ToUpper(strSlice[2])
	aArab, err1 := strconv.Atoi(a)
	cArab, err2 := strconv.Atoi(c)
	if (err2 != nil && err1 == nil) || (err1 != nil && err2 == nil) {
		fmt.Println("Некорректные данные(разные системы исчисления)")
		return
	}
	if b != "+" && b != "-" && b != "/" && b != "*" {
		fmt.Println("Некорректные данные(неверный знак)")
		return
	}
	//разделение на 2 потока. 1 - арабские, 2 - римские
	switch {
	case err1 == nil:
		fmt.Println(err1, err2)
		if aArab > 10 || aArab < 1 || cArab > 10 || cArab < 1 {
			fmt.Println("Некорректные данные(веедены недопустимые данные)")
			return
		}
		if errors.Is(err2, err1) && err1 == nil {
			switch b {
			case "+":
				fmt.Println(aArab + cArab)
				return
			case "-":
				fmt.Println(aArab - cArab)
				return
			case "*":
				fmt.Println(aArab * cArab)
				return
			case "/":
				fmt.Println(aArab / cArab)
				return
			}
		}
	case err1 != nil:
		var aRome int
		var cRome int
		var result int
		for i, rom := range rome {
			if rom == a {
				aRome = i
			}
			if rom == c {
				cRome = i
			}
		}
		if aRome > 10 || aRome < 1 || cRome > 10 || cRome < 1 {
			fmt.Println("Некорректные данные(веедены недопустимые данные)")
			return
		}
		switch b {
		case "+":
			result = aRome + cRome
		case "-":
			result = aRome - cRome
		case "*":
			result = aRome * cRome
		case "/":
			result = aRome / cRome
		}
		if result < 0 {
			fmt.Println("Некорректный результат (отрицательный римский результат)")
			return
		}
		for i, rom := range rome {
			if result == i {
				fmt.Println(rom)
			}

		}
	}
}
