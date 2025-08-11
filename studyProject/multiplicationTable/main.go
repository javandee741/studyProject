package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Программа выводит таблицу умножения для введённого числа.")
	fmt.Println("Для выхода введите 0.\n")

	for {
		fmt.Print("Введите число (0 для выхода): ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // Удаляем символ новой строки (\n)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите целое число!")
			continue // Возвращаемся к запросу числа
		}

		if num == 0 {
			fmt.Println("Выход из программы.")
			break
		}

		printTable(num)
		fmt.Println() // Пустая строка для разделения таблиц
	}
}
