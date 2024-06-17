package main

// Всегда название файла начинается с объявления с пакета. В любом проекте должен быть пакет main - иначе запуска приложения не будет.

// Определяет название пакета, чтобы иметь обращение к этому файлу.

import (
	"errors"
	"fmt"
	"math"
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

func panicDeferRecover() {
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
	fmt.Println("panicDeferRecover()")
}

func mapUsage() {
	// map - Это словарь.
	users := map[int]string{
		1: "Me",
		2: "You",
		3: "She",
		4: "He",
	}

	fmt.Println(users)
	fmt.Println(users[1])
	// В случае получения по не существующему ключу - выдаст дефолтное значение для типа value(значение). В данном случае - пустая строка.
	// Никакой ошибки в этом нет, по сравнению с Python. Аналогия get - но там вернется Null.
	fmt.Println(users[6])

	// Так мы точно можем узнать, существует ли такой ключ. Возвращается bool значение.
	who, exist := users[6]

	if !exist {
		fmt.Println("Такого нет")
	} else {
		fmt.Println(who)
	}

	for key, value := range users {
		// Получаем ключ значение (мы не проходимся, тк данные не упорядоченны), порядок может быть случайным.
		fmt.Println(key, value)
	}

	fmt.Println()
	// Добавление в словарь(map) нового значения.
	users[6] = "idk"
	// Изменение значения по ключу.
	users[1] = "not me"
	// Удаление из словаря значения по ключу.
	delete(users, 2)

	for key, value := range users {
		fmt.Println(key, value)
	}

	// Объявляем словарь, но не инициализируем его.
	var users2 map[int]string
	// В этом месте произойдет ошибка(паника), тк словарь не инициализирован.
	// users2[1] = "Me"

	// Инициализируем словарь ключевым словом make, теперь можем добавлять/изменять значения.
	users2 = make(map[int]string)

	users2[1] = "me"

	fmt.Println(users2)

	fmt.Println(len(users))
	// fmt.Println(cap(users)) - У словаря нет вместимости, только длина.
}

func structSimple() {
	// В таком случае мы инициализировали структуру, и уже не может её переобределить.
	user := struct {
		name   string
		age    int
		sex    string
		height int
	}{
		"Yaroslav", 20, "Male", 180,
	}

	fmt.Println(user)
}

func structInit() {
	type User struct {
		name   string
		age    int
		sex    string
		height int
	}

	user := User{"Yaroslav", 20, "Male", 180}

	fmt.Println(user)

	fmt.Println(user.age, user.name)
}

type UserBase struct {
	// Создаем структуру User.
	name   string
	age    int
	sex    string
	height int
}

func (us UserBase) printInfoUser() {
	// Это метод структуры. В данном случае это ресивер метод (может по значению или по ссылке на структуру).
	// В данном случае по ссылке.(создаётся копия, структура не меняется)
	fmt.Println(us.age, us.name, us.height, us.sex)
}

func (us *UserBase) changeNameUser(name string) {
	us.name = name
}

func NewUser(name, sex string, age, height int) UserBase {
	// Конструктор для User.
	return UserBase{
		name:   name,
		sex:    sex,
		age:    age,
		height: height,
	}
}

func printInformationUser(user UserBase) {
	// Возвращает инфо по полям.
	fmt.Println(user.age, user.name, user.height, user.sex)
}

func userManipulate() {
	user1 := NewUser("Petya", "Male", 30, 160)
	printInformationUser(NewUser("Yaroslav", "Male", 20, 180))
	printInformationUser(user1)

	user1.printInfoUser()
	user1.changeNameUser("Kostya")
	user1.printInfoUser()
}

type Age int

func (a Age) validateAge() bool {
	return a > 18
}

func ageValidate() {
	age := Age(14)

	fmt.Println(age, age.validateAge())
}

type Shape interface {
	ShapeArea
	ShapePerimetr
}

type ShapeArea interface {
	Area() float32
}

type ShapePerimetr interface {
	Perimetr() float32
}

type Circle struct {
	radius float32
}

type Square struct {
	side float32
}

func (squre Square) Area() float32 {
	// Возвращаем площадь для квадрата.
	return squre.side * squre.side
}

func (square Square) Perimetr() float32 {
	return square.side * 4
}

func (circle Circle) Area() float32 {
	// Возвращаем площадь для круга.
	return circle.radius * circle.radius * math.Pi
}

func (circle Circle) Perimetr() float32 {
	return circle.radius * math.Pi * 2
}

func showShapeArea(shape Shape) float32 {
	// Вызываем метод Area у интерфейса.
	return shape.Area()
}
func showShapePerimetr(shape Shape) float32 {
	return shape.Perimetr()
}

func initShapes() {
	// Инициализируем переменные и вызываем метод, который принимает интерфейс Shape, содержащий метод Area.
	square := Square{3}
	circle := Circle{1.7}

	fmt.Println(showShapeArea(square))
	fmt.Println(showShapeArea(circle))

	fmt.Println(showShapePerimetr(square))
	fmt.Println(showShapePerimetr(circle))

	// Передаем инициализированную структуру и тк она удовлетворяте пустому интерфейсу возвращаем все дефолтные поля.
	printInterface(UserBase{})
	printInterface(Square{5})
}

func printInterface(i interface{}) {
	// Мы передаем пустой интерфейс, поэтому нам не важно какой тип мы передаем, тк любой тип удовлетворяет пустому интерфейсу.
	// fmt.Printf("%+v\n", i)

	switch value := i.(type) {
	case int:
		fmt.Println(value, "- int")
	}

	str, ok := i.(string)
	if ok {
		fmt.Println(len(str))
	} else {
		fmt.Println("This interface -", i, "has no len")
	}
}

// Точка входа в приложение.
func main() {
	printInterface("qwe")
	fmt.Println()
	printInterface(123)

	fmt.Println()
	fmt.Println()

	initShapes()
}
