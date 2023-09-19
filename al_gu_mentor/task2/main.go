package main

import (
	"errors"
	"fmt"
)

func main() {

	err := calculate()
	if err != nil {
		fmt.Println(err)
	}

}

func calculate() error {
	var userNum int

	fmt.Println("Введите 0 для ошибки деления на 0, или букву для других ошибок")
	_, err := fmt.Scan(&userNum)
	if err != nil {
		return errors.New("test error 1")
	}

	_, err2 := fmt.Println(10 / userNum)
	if err2 != nil {
		return errors.New("test error 2")
	}

	return nil
}
