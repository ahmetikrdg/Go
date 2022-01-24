package main

import (
	"fmt"
)

func main() {
	x := 5
	fmt.Println(x) //5
	dobule(x)      //10
	fmt.Println(x) // 5 yazdırdı
	//değişmeme nedeni şu go oass by value şeklinde yani değer geçen bir prog dili.
	//bir fonksiyon çalıştırıldığı zaman aldığı parametre agrümanın kopyasıdır.
	//5 i yazdırdı sonra metodu çalıştırdım dobule() fonksiyonumu çağırdım double fonksiyonumun
	//argümanı ne x. x ne yerine geçiyo doubledeki (num int) parametresi yerine geçiyo
	//x 5 num'da 5 oluyo ancak num 5 olurken x değeri numa gelmiyo x'in kendisi değil kopyası gelir
	//sonra 3. cüde tekrar x yazdırınca 5 geliyo neden çünkü double giden o x'in kopyasıdır o nedenle x  değişmez
}

func dobule(num int) {
	num *= 2
	fmt.Println(num)
}
