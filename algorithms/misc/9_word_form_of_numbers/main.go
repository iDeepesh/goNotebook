package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	stop := false
	for !stop {
		fmt.Print("Please enter a number smaller than 1 Billion or (q) to quit: ")
		s.Scan()
		t := s.Text()
		if t == "q" {
			stop = true
		}

		n, _ := strconv.Atoi(t)
		fmt.Println(getWordForm(n))
	}
}

func getWordForm(n int) string {
	d, do, m := 1000, 0, 1000
	var s string
	tn := n % d
	if tn > 0 {
		s = getHundreds(tn)
	}

	d, do = d*m, d
	tn = n % d
	tn = tn / do
	if tn > 0 {
		s = getHundreds(tn) + " thousand " + s
	}

	d, do = d*m, d
	tn = n % d
	tn = tn / do
	if tn > 0 {
		s = getHundreds(tn) + " million " + s
	}

	return s
}

func getHundreds(n int) string {
	tn := n / 100
	var s string
	if tn > 0 {
		s = getDigit(tn) + " hundred "
	}

	tn = n % 100
	if tn > 0 {
		s = s + getTens(tn)
	}

	return s
}

func getTens(n int) string {
	tn := n / 10
	if tn == 1 {
		return getTeen(n)
	}

	s := getDigit(n % 10)

	if tn > 1 {
		s = getTies(tn) + " " + s
	}

	return s
}

func getTeen(n int) string {
	var s string
	switch n {
	case 10:
		s = "ten"
	case 11:
		s = "eleven"
	case 12:
		s = "twelve"
	case 13:
		s = "thirteen"
	case 14:
		s = "forteen"
	case 15:
		s = "fifteen"
	case 16:
		s = "sixteen"
	case 17:
		s = "seventeen"
	case 18:
		s = "eighteen"
	case 19:
		s = "ninteen"
	}

	return s
}

func getTies(n int) string {
	var s string
	switch n {
	case 2:
		s = "twenty"
	case 3:
		s = "thirty"
	case 4:
		s = "forty"
	case 5:
		s = "fifty"
	case 6:
		s = "sixty"
	case 7:
		s = "seventy"
	case 8:
		s = "eighty"
	case 9:
		s = "ninty"
	}

	return s
}

func getDigit(n int) string {
	var s string
	switch n {
	case 1:
		s = "one"
	case 2:
		s = "two"
	case 3:
		s = "three"
	case 4:
		s = "four"
	case 5:
		s = "five"
	case 6:
		s = "six"
	case 7:
		s = "seven"
	case 8:
		s = "eight"
	case 9:
		s = "nine"
	}

	return s
}
