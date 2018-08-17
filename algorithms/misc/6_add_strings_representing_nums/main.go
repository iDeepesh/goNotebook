package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter the number of digits in first string: ")
	s.Scan()
	n1, _ := strconv.Atoi(s.Text())

	fmt.Print("Please enter the number of digits in second string: ")
	s.Scan()
	n2, _ := strconv.Atoi(s.Text())

	s1 := generateStringForNumber(n1)
	fmt.Println("First string:", s1)
	s2 := generateStringForNumber(n2)
	fmt.Println("Second string:", s2)
	sum := addNumberStrings(s1, s2)
	fmt.Println("Sum:", sum)
}

func addNumberStrings(s1, s2 string) string {
	r1, r2 := []rune(s1), []rune(s2)

	l := len(s1)
	if l < len(s2) {
		l = len(s2)
	}
	r := make([]rune, l+1)

	c := 0
	i, j, k := len(r1)-1, len(r2)-1, len(r)-1
	for {
		if i < 0 && j < 0 {
			break
		}

		sum := c
		if i >= 0 {
			sum += int(r1[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(r2[j] - '0')
			j--
		}

		c = sum / 10
		sum = sum % 10
		r[k] = rune(sum + '0')
		k--
	}

	if c > 0 {
		r[k] = rune(c + '0')
	} else {
		r = r[1:]
	}

	return string(r)
}

func generateStringForNumber(l int) string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var sb strings.Builder
	for i := 0; i < l; i++ {
		sb.WriteString(strconv.Itoa(r.Intn(9)))
	}

	return sb.String()
}
