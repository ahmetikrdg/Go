package main

import "fmt"

var sayi int = 4 //func dışında olduğum için sayi:=4 gibi bir kullanım yapamam

func main() {
	/*
				var message string
				message = "Merhaba Go!"
		  	var message string = "Go"
				var message="Merhaba Go!"

				var a, b, c int = 3, 4, 5
				var a, b, c = 3, true, 4.15

				fmt.Println(a, b, c)

					var a int
					var b string
					var c float32
					var d bool

					u := 55

					v, n := "ali", "veli" //gidecek verinin tipi biliniyorsa vermekte fayda var

	*/

	var (
		degisken1 int = 1
		degisken2     = "Veli"
	)
	fmt.Println(degisken1, degisken2)

}
