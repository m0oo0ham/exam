package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("EXAMPLE : uncopyrightable,4")
	fmt.Print("INPUT : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	argument := strings.Split(input, ",")
	arg1 := argument[0]
	arg2, _ := strconv.Atoi(argument[1])
	solution(arg1, arg2)

	// arg1 := "uncopyrightable"
	// arg2 := 4
	// solution(arg1, arg2)
}

func solution(s string, n int) {
	index := map[int]bool{}
	index[0] = true
	for i := 1; i < len(s); i++ {
		if i%(n-1) == 0 {
			index[i] = true
		}
	}

	rows := []string{}
	for i := 0; i < n; i++ {
		text := strings.Repeat(" ", len(s))
		rows = append(rows, text)
	}

	slideDown := true
	k := 0
	for i := 0; i < len(s); i++ {

		_, exist := index[i]
		if exist {
			if i != 0 {
				if slideDown {
					slideDown = false
				} else {
					slideDown = true
				}
			}

		}
		rows[k] = rows[k][:i] + string(s[i]) + rows[k][i+1:]
		if slideDown {
			k++
		} else {
			k--
		}
	}
	fmt.Println("OUTPUT : ")
	for _, v := range rows {
		fmt.Println(v)
	}
}
