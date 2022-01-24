package main

import (
	"fmt"
)

func main() {
	/*
		var myString = "1"
		fmt.Println(myString)
		//string convert to int
		number, _ := strconv.Atoi(myString) // _ boş tanımlayıcıdır sebebi ise hata nesnesi dönerse onu almayayım boş tanımlayıcı koydum
		fmt.Println(number)

		result := number + 2
		fmt.Println(result)
	*/
	//CASTING

	var i int = 55
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(u)

}
