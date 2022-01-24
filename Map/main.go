package main

import "fmt"

func main() {
	myMap := make(map[int]string) //mapin anahtarının tipi int karılığı string
	myMap[43] = "Foo"
	myMap[12] = "Bar"
	fmt.Println(myMap)

	//state
	states := make(map[string]string)
	states["IST"] = "İstanbul"
	states["ANK"] = "Ankara"
	states["ANT"] = "Antalya"
	fmt.Println(states)

	//Şehir listesinde ank anahtar adına sahip veriyi elde et
	antalya := states["ANT"]
	fmt.Println("Seçili Şehir: ", antalya)

	//ANK anahtar adına sahip veriyi sil
	delete(states, "ANK")
	fmt.Println(states)

	//Ekleme
	states["ERZ"] = "Erzurum"

	for k, v := range states {
		fmt.Printf("%v: %v \n", k, v)
	}

	//states map nesnesinin elemanı adedince kapasitesine sahip bir key listesi oluştur

	keys := make([]string, len((states)))
	i := 0

	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
}
