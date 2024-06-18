package main

type metr uint
type kilometr uint

type metr2 = uint
type kilometr2 = uint

// В таком случае когда мы присваеваем один тип другому определенному типу, то разные типы с одним и тем же исходным - будут равны и иметь исходный тип.

func testingType() {

	var metr_to_object metr = 12
	var km_to_obj kilometr = 1
	var distance_object uint = 10

	returnDistance(metr_to_object)
	// returnDistance(km_to_obj)
	// returnDistance(distance_object)
	// Вызов этих двух методов будет не верным, тк мы создали отдельный тип, который хранится в памяти, и даже то что они имеют общий начальный тип
	// Не означает, что это одинаковый объект.
	returnDistance(metr(km_to_obj))
	returnDistance(metr(distance_object))
	// В таком случае сработает, тк мы привели тип к необходимому.
}

func returnDistance(metrs metr) {
	println("Distance to object - ", metrs)
}

func testingType2() {

	var metr_to_object metr2 = 12
	var km_to_obj kilometr2 = 1
	var distance_object uint = 10

	returnDistance2(metr_to_object)
	returnDistance2(km_to_obj)
	returnDistance2(distance_object)
	// Вызов этих двух методов будет не верным, тк мы создали отдельный тип, который хранится в памяти, и даже то что они имеют общий начальный тип
	// Не означает, что это одинаковый объект.
}

func returnDistance2(distance metr2) {
	println("Distance to object 2 - ", distance)
}
