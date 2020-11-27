package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("EXAMPLE : Hello world")
	fmt.Print("INPUT : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	solution(input)

	// input := "I thought I was a good officer back then but now when taking a step back and seeing other officers doing their work I know I was never a really good cop"
	// solution(input)
}

func solution(s string) {
	words := map[string]int{}
	text := strings.Split(s, " ")
	for _, v := range text {
		words[v]++
	}
	fmt.Println("OUTPUT :")
	for i, v := range words {
		if v > 1 {
			fmt.Printf("\"%v\" - %v times\n", i, v)
		}
	}
}
