package main

import (
	"fmt"
)

func main() {
	myArray := [3]int{1, 10, 100}
	fmt.Println(myArray) //1 10 100
	dobule(myArray)      //2 20 200
	fmt.Println(myArray) //1 10 100

	//pass by value go olduğu için myArray 1 10 100 değerini verid peki neden slicelerdeki gibi olmadı
	//slicelere kopyasını gönderiyoruzya kopyasını gönderdiğimiz arrayler birbirinden bağımsız sliceler gibi aynı underline arrayi tutmuyorlar
	//farklı underlade arrayleri var bunların. birisini değiştirince orjinal değer sabit kalır.

}

func dobule(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2 //fonksiyon slicenin kopyasını değiştirdi
	}
	fmt.Println(arr)
}
