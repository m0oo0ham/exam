package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("EXAMPLE : \"1 2 3\",\"4 5 6\",\"7 8 9\"")
	fmt.Print("INPUT : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	inputSets := strings.Split(input, ",")
	for i := 0; i < len(inputSets); i++ {
		inputSets[i] = strings.Trim(inputSets[i], "\"")
	}
	arr := [][]int{}
	for i := 0; i < len(inputSets); i++ {
		intSlice := []int{}
		for j := 0; j < len(inputSets); j++ {
			splitText := strings.Split(inputSets[i], " ")
			stringToInt, _ := strconv.Atoi(splitText[j])
			intSlice = append(intSlice, stringToInt)
		}
		arr = append(arr, intSlice)
	}
	solution(arr)

	// arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// solution(arr)
}

func solution(arr [][]int) {
	sum := 0
	k := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		value := arr[i][i]
		sum += value
	}
	for i := 0; i < len(arr); i++ {
		value := arr[i][k-i]
		sum += value
	}
	if len(arr)%2 == 1 {
		i := len(arr) / 2
		sum -= arr[i][i]
	}
	fmt.Println("OUTPUT : ", sum)
}
