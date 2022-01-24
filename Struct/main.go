package main

import "fmt"

//structlar yani yapılar alanlardan oluşan veri topluluğudur. Alanlar field olarak geçer.
//Diğer prog dillerinde structun karşılığı olan şey sınıflardır class.
//Go prog dilinde sınıf yoktur yada class kavramı yoktur.
//bunların içerisinde alanlara değişkenlerede field denir. sınıfların içindeki değişkenelerede field denir.
//type bildirimi struct oluşturmaya yarar. class person struct {} gibi c# da ama burada type person strucu{}

func main() {
	//Nesne Örneği Oluşturma
	/*
			x := Human{FirstName: "Ahmet", LastName: "Karadağ", Age: 22}
			fmt.Println(x.FirstName)


		x := new(Human)
		x.FirstName = "Ali"
		fmt.Println(x.FirstName)
	*/

	//ctor kullanma
	/*
		x := NewHuman()
		x.Age = 22
	*/

	//parametreli ctor kullanımı
	x := NewHumanParams("Ahmet", "Karadağ", 22)
	fmt.Println(x.Age)
}

type Human struct { //struct tanımlama
	FirstName string
	LastName  string
	Age       int
}

//ctor oluşturma işlemi default ve boş yapıcı metot
func NewHuman() *Human { //Human döneceğim diyorum
	h := new(Human)
	return h
}

/*
/parametreli ctor oluşturma
func NewHuman(firstname, lastname string, age int) *Human {
	h := new(Human)
	h.FirstName = firstname
	h.LastName = lastname
	h.Age = age
}//bunu yapmayız golang metot overloadingi desteklemez. doyki sen bunu farklı bir isimli farklı bir metot olarak tanımlana gerek
*/

func NewHumanParams(firstname, lastname string, age int) *Human { //farklı isimle oluşturduk
	h := new(Human)
	h.FirstName = firstname
	h.LastName = lastname
	h.Age = age
	return h
}
