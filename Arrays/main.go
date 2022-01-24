package main

import "fmt"

func main() {
	//diziler

	myArray1 := [3]string{}
	myArray1[0] = "Ali"
	myArray1[1] = "Veli"
	myArray1[2] = "Selami"
	fmt.Println(myArray1)

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"
	println(colors[1])
	colors[1] = "Bordo"
	println(colors[1])

	var numbers = [5]int{4, 5, 6, 7, 3}
	fmt.Println(numbers)
	fmt.Println("Number of numbers", len(numbers))

	myArray2 := [...]int{4, 5, 6, 10, 100, 15, 60, 2, 58, 5, 4, 5, 8} //otomatic size yeteniğini kullandım istediğim kadar sayı gönderebilirim
	fmt.Println(myArray2)

	var cars [3]string
	cars[0] = "Fiat"
	cars[1] = "Kia"

	i := 0
	for i <= len(cars)-1 { //cars 0 dan başlayacağı için -1 dedik
		fmt.Println(cars[i])
		i++
	}

	for i = 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}

	for i, value := range cars {
		fmt.Println(i, "-- ", value)
		//cars[i] --> buda olur ama valueyi kaldırmam gerek
		//veya for i, value yerine _, valuede yapılabilir
	}
}
