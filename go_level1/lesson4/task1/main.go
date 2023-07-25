package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите цифры: ")
	inputNums := []int64{}

	for scanner.Scan() {
		if len(scanner.Text()) != 0 {
			num, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
			inputNums = append(inputNums, num)
		} else {
			break
		}

	}

	for i := 1; i < len(inputNums); i++ {
		for j := i; j >= 1 && inputNums[j] < inputNums[j-1]; j-- {
			inputNums[j], inputNums[j-1] = inputNums[j-1], inputNums[j]
		}
	}
	fmt.Print("Результат сортировки: ", inputNums)
}
