package main

import "fmt"

//C# daki slice go'dki liste tekamül eder
//sliceler arrraylerin biraz daha genişletilmiş versiyonu. arraylar oluşturulduğunda eleman sayısını bilirken
//sliceler oluşturulduğunda eleman sayısını bilmiyor
//https://www.w3schools.com/go/go_slices.php
func main() {

	myslc := []int{1, 2, 3} //söz dizimi array gibi ama biz söz dizimi yaparken [] burada eleman sayısı belirtmedik. Çünkü sliceler oluşturulurken eleman sayısını bilmezler
	fmt.Println(myslc)

	//array ve slice farkını net görelim
	fmt.Println("Array ve Slice farkı")

	var myArr [4]int
	fmt.Println(myArr)
	myArr[0] = 4
	fmt.Println(myArr) //4 0 0 0 döner

	var mySlc2 []int
	fmt.Println(mySlc2) //boş bir array parçası boş bir slice döner []
	//	mySlc2[0] = 10      //hata verir
	//fmt.Println(mySlc2) //uzunluğu geçtin der. burada slice belirtiliyo ama oluşturulmamış durumda

	mySlc2 = make([]int, 4)
	fmt.Println(mySlc2) //aynen arraylerde oludğu gibi make ile oluşturdum

	//go kodları comlier ederken arrayler daha güvenilirdir. neden çünkü biz arrayin index sınırını biliriz ve bu complier timeda değiştirilmez bu sebeple daha güvenilir ve daha hızlı.
	//bir arrayın kopyasını aldığımız zaman arrayın kendisini kopyalarız ama slicelerda bu durum geçerli değil.
	//arrayler diğer prog dillerinde olduğu gibi referanslarını değil kendi değerlerini kopyalıyo

	myArr3 := [3]int{1, 2, 3}
	fmt.Println(myArr3)
	myArr4 := myArr3
	fmt.Println(myArr4) //1 2 3 yazdırdı süper
	myArr3[0] = 100
	fmt.Println(myArr3) //burada 100 2 3
	fmt.Println(myArr4) // burada ise 1 2 3
	//görüldüğü gibi myarr3 değişirken myarr4 değişmedi çünkü referansını almıyo direk kendisini kopyalıyo
	//fakat slice'de böyle değil

	mySlcc := []int{1, 2, 3}
	fmt.Println(mySlcc)
	mySlcc2 := mySlcc
	fmt.Println(mySlcc2)
	mySlcc2[0] = 33
	fmt.Println(mySlcc2)
	fmt.Println(mySlcc) //33 2 3 olarak gelir ikisindede

	//sliceler pass by referance olarak işler
	//arrayler pass by referance olarak işler

	mySlcNew := []int{1, 2, 3}
	fmt.Println(mySlcNew)

	mySlcNew = append(mySlcNew, 10, 15)
	fmt.Println(mySlcNew)
}
