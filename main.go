package main

// Всегда название файла начинается с объявления с пакета. В любом проекте должен быть пакет main - иначе запуска приложения не будет.

// Определяет название пакета, чтобы иметь обращение к этому файлу.

import (
	"errors"
	"fmt"
)

// Импортируем стандартную бибилиотеку.

/*
Таким способом делаю многострочный комментарий.
Вот тут вторая строка.
А тут третья.
*/

// Точка входа в приложение.
func main() {

	minimal, min_key := findMin(1, 3, 5, 0)

	fmt.Println(minimal, min_key)
}

func test() {

	var message string

	// Строка в go всегда записываается в двойных кавычках "".
	// Руна (rune) записывается в одинарных кавычках - ''. '' - запись символа.
	var number int
	var boolean bool

	// rune / byte - это замещение символов в строку или значение и наоборот.

	boolean = false

	fmt.Println(message, number, boolean)
}

func print(message string) (int, error) {
	// Фнункция, которая выводит в терминал сообщение и возвращает два значения - длину строки и наличие ошибок при выполнении.
	return fmt.Println("Функция print()", message)
}

func returnMessage(message string) string {
	// Функция возвращающая конкатенацию строки.
	return "Ваше сообщение - " + message
}

func returnFormatString(message string, age int8) string {
	// Функция форматирования строки, исходя из параметров переданных в функцию.
	result := fmt.Sprintf("Привет, %s! Тебе %d лет.", message, age)
	return result
}

func returnErrors(age int) (string, error) {
	// По конвенции error должен возвращаться последним значением.

	if age >= 18 && age < 20 {
		return "Вам больше 18, но нет 20", nil
	} else if age < 18 {
		return "Вам нет 18", errors.New("oops..")
	}

	// Ошибки должны быть на английском языке и начинаться с маленькой буквы, быть короткими.

	return "Вам больше 18", errors.New("hmm...")
}

func findMin(numbers ...int) (int, int) {
	if len(numbers) == 0 {
		return 0, 0
	}

	min := numbers[0]
	min_key := 0

	for key, value := range numbers {
		if value < min {
			min = value
			min_key = key
		}

	}

	return min, min_key
}
