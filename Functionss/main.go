package main

import "fmt"

func main() {
	fmt.Println(add(10, 15))
	message := "hi"
	sayHello(&message) //pointer gönderiyorum. mesajın hafızdaki adresini
}

func add(x int, y int) int {
	return x + y
}

func sayHello(message *string) { //poniter istiyorum diyor
	println(*message)
}

//1300
