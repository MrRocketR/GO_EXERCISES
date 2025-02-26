package _kyu

import "strings"

var numbers map[int]string = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

//1 and 3999  MMM CMXC IX 3000 + 500 +
//2008 is written as 2000=MM, 8=VIII; or MMVIII.
// 1666 uses each Roman symbol in descending order: M DC LXVI.
// 1948 = MCMXLVIII
//Тысяча Девятьсот Сорок Восемь = M CM XL VIII

func Solution(number int) string {
	var roman string

	M := number / 1000 //3
	//тысячи 1988
	if M > 0 {
		for i := 1; i <= M; i++ {
			roman += numbers[1000]
		}
		number = number - (M * 1000)
	}
	C := number / 100 //999 / 100 = 9

	//сотки
	if C > 0 {
		if C < 4 {
			for i := 1; i <= C; i++ {
				roman += numbers[100]
			}
		} else if C == 4 {
			roman += numbers[100]
			roman += numbers[500]
		} else if C == 5 {
			roman += numbers[500]
		} else if C < 9 { //8
			roman += numbers[500]
			for i := 6; i <= C; i++ {
				roman += numbers[100]
			}
		} else if C == 9 {
			roman += numbers[100]
			roman += numbers[1000]
		}
		number = number - (C * 100)
	}
	X := number / 10 //99 / 10 = 9

	//десятки
	if X > 0 {
		if X < 4 {
			for i := 1; i <= X; i++ {
				roman += numbers[10]
			}
		} else if X == 4 {
			roman += numbers[10]
			roman += numbers[50]
		} else if X == 5 {
			roman += numbers[50]
		} else if X < 9 { //8
			roman += numbers[50]
			for i := 6; i <= X; i++ {
				roman += numbers[10]
			}
		} else if X == 9 {
			roman += numbers[10]
			roman += numbers[100]
		}
		number = number - (X * 10)
	}

	if number > 0 {
		if number < 4 {
			for i := 1; i <= number; i++ {
				roman += numbers[1]
			}
		} else if number == 4 {
			roman += numbers[1]
			roman += numbers[5]
		} else if number == 5 {
			roman += numbers[5]
		} else if number < 9 { //8
			roman += numbers[5]
			for i := 6; i <= number; i++ {
				roman += numbers[1]
			}
		} else if number == 9 {
			roman += numbers[1]
			roman += numbers[10]
		}
	}

	return roman
	// convert the number to a roman numeral
}

func Solution2(n int) string {
	if n <= 0 {
		return ""
	}
	var sb strings.Builder
	for i, v := range []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} {
		for n >= v {
			n -= v
			sb.WriteString([]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}[i])
		}
	}
	return sb.String()
}
