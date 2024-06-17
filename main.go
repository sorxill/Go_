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

var msg string

func init() {
	// Функция инициализации пакета. Вызывается до main.
	msg = "Я из инициализации пакета."
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

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func printMessage(message string) {
	message += " (из printMessage())"
	fmt.Println(message)
}

func changeMesage(message *string) {
	// Чтобы изменить указатель необходимо указать '*' перед переменной. в аргументе перед типом.
	*message += " Я тут изменил немного текст)"
}

func pointers() {
	number := 5

	var p *int

	p = &number
	fmt.Println(*p)
	fmt.Println(number)
	*p = 10

	fmt.Println(number)
	fmt.Println(*p)
}

func slicesAndLists() []int {
	/* Slice - обертка над массивами в Go. Он не имеет строгой длины, в отличии от массива.

	!!! slice внутри себя хранит ссылку на массив, который можно передать в функцию, при изменении внутри функции будет меняться и аргумент.

	msg := []int{1, 2, 3} - slice без указания длины.
	msg1 := [3]int{1, 2, 3} - massive с указанием длины.
	*/
	integers := []int{1, 2, 3}
	return integers
}

func matrixManipulate() {
	// Матрицы по сути двумерные массивы/слайсы.
	var matrix [][]int = make([][]int, 10)

	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			if y == x {
				matrix[y][x] = x
			}
		}
	}

	for _, value := range matrix {
		// Перебор индекса в слайсе и его значения.
		fmt.Println(value)
	}

	for key, value := range matrix {
		// Из первого преобразования матрицы выбираем те места, которые не нулевые.
		for index, x := range value {
			if key == x && index == key {
				fmt.Println(key, x)
			}
		}
	}

}

func alwaysFor() {
	// Бесконечный цикл - аналогия While в Python.
	var counter int = 0

	for {
		if counter == 10 {
			break
		}
		counter++
		fmt.Println(counter)
	}
}

func panicFunc(a int, b int) int {

	if b == 0 {
		panic("division by zero!")
	}

	return a / b
}

func recoveryFunction() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		// В случае, если у нас вызвана паника и сообщение не nil - мы вернем сообщение об ошибке.
		fmt.Println("panic message - ", recoveryMessage)
	}
	// Если паники не было, то выполнится только этот код. (В случае паники в данной ситуации это сообщение тоже выполнится.)
	fmt.Println("This is recover Function")
}

// Точка входа в приложение.
func main() {
	/* Чем выше defer находится, тем позже эта функция будет  в коде. В случае срабатывания panic выполнение программы будет прекращено незамедлительно.
	Но в случае если у нас есть recover - программа не будет прекращена незамедлительно, управление будет передан main потоку.
	*/

	// defer recoveryFunction()

	// fmt.Println(panicFunc(5, 1))

	// printMessage("test")
	// fmt.Println("main()")

	defer recoveryFunction()

	// вызываем панику делением на нуль.
	fmt.Println(panicFunc(5, 0))

	// В данном случае эти две функции не будут вызваны, тк была вызвана паника.
	printMessage("test")
	fmt.Println("main()")
}
