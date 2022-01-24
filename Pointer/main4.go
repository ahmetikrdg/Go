package main

import (
	"fmt"
)

func main() {
	x := 5
	fmt.Println(x) //5

	//pointerları neden kullanırız.
	//bir değeri olduğu yerde değiştirebilmek için kullanırız avantajı örn
	//verimiz çok büyük arrayden veya sliceden bunun işlem yaparken kopyasını almak efektif sonuç değildir.
	//kopyası yerine olduğu yerde değiştirmek yani pointerını almak bizim açımızdan daha avantajlı olur.
	//pointerı bu nedenle kullanıyoruz. bir veriyi bulunduğu yerde değiştirmek için kullanırız

	dobule(&x)     //10 ampersan & operatörü koydum ve yani artık bir pointerı parametre olarak gönderebiliyorum
	fmt.Println(x) //10
	//artık x'in değeri 10 bunu pointer yardımıyla becerdik. x'i prametre olarak göndermek yerine adresini gönderdik.
	//x'in adresindeki değeri 2 ile çarptık. Buda x değerinin değişmesini sağlamış oldu.
	//biz go'da değerin kendisini olduğu yerde değiştirmek için pointer kullanıyoruz.

}

func dobule(num *int) { //x'in adresini alacağım. nasıl belirtirdik poniterla. poinrterı veri tipine dönüştürmek için asterix kullanırım *int. pointer veri tipiyle gönderileni yakaladık
	fmt.Println(num) //num artık x in adresini tutuyo
	//numun içindeki değeri nasıl yakalıyodum derefer ediyodum başına asterix koyuyodum.
	*num = *num * 2   //dönderdiğimin iki katını alıyorum
	fmt.Println(*num) //adresteki değeri yazdırsın istiyorum

	//5'in kendisini göndermedik adresini gönderdik. adres'i num *int ile yakaladık. adresteki değeri iki katına çıkardım. o değeri yazdırdım

} //double içinde artık pointer method deriz.
