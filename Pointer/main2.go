package main

import (
	"fmt"
)

func main() {
	mySlc := []int{1, 10, 100}
	fmt.Println(mySlc) //1 10 100
	dobule(mySlc)      //2 20 200
	fmt.Println(mySlc) //2 20 200

	//sliceler pass by referansdır peki orada durumlar nasıl olur?
	//go yine aynı şekilde değeri geçer ama buradaki olay slicelerin yapısıyla ilgili
	// double(mySlc) ile slicenin bir kopyasını fonksiyonumuza gönderdik.
	//donksiyon slicenin kopyasını değiştirdi. pkei neden orjinalide değişti
	//çünkü slicelerin özelliği neydi arka planda bunların pointerleri vardı aynı underline array tutuyorlardı.
	//bir slice değişince underline array değişir. o array değişince sliceıda değiştirir.
	//aslında referans değiştirmiyo değerini değiştiriyo.
	//özetle

	//dobule(mySlc) bunu gönderince argümanın kendisi değil kopyası gider. o kopyayı değiştirdiği için hem orjinal hem kopya slice değişiyo biz bunu referans değişiyo diye aldırıyoruz ama değişen referansın kendisi.
}

func dobule(slc []int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2 //fonksiyon slicenin kopyasını değiştirdi
	}
	fmt.Println(slc)
}

//bu örneği arraylerle yapalım main3.go'da
