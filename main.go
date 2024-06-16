package main

// Всегда название файла начинается с объявления с пакета. В любом проекте должен быть пакет main - иначе запуска приложения не будет.

// Определяет название пакета, чтобы иметь обращение к этому файлу.

import (
	"fmt"
	"reflect"
)

// Импортируем стандартную бибилиотеку.

/*
Таким способом делаю многострочный комментарий.
Вот тут вторая строка.
А тут третья.
*/

// Точка входа в приложение.
func main() {
	// test()
	var message string = "test"

	n, err := print(message)
	fmt.Println(n, err)

	result_mesage := returnMessage("returned message")
	fmt.Println(result_mesage)

	f_string_message := returnFormatString("Ярослав", 20)
	fmt.Println(f_string_message)
}

func test() {
	var test int8 = 123

	var message string

	// Строка в go всегда записываается в двойных кавычках "".
	// Руна (rune) записывается в одинарных кавычках - ''. '' - запись символа.
	var number int
	var boolean bool

	// rune / byte - это замещение символов в строку или значение и наоборот.

	boolean = false

	fmt.Println(message, reflect.TypeOf(message), len(message))
	fmt.Println(test, message, number, boolean)
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
