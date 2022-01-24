package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	number := 45

	fmt.Printf("Gelen Sayi: %v\n", number) //v ile gelen değeri oraya atıyorum

	//konsole veri girişi: bunun için bunfio kullanırım
	reader := bufio.NewReader(os.Stdin) //os yani işledim sistemi içindeki stdin nesnesini kullan dışarıdan gelen akışı yakalayarak uygulama içinde kullanır
	fmt.Print("Veri Giriniz: ")
	str, _ := reader.ReadString('\n') //yukarıda oluşturduğun readerden veriyi bekliyorum. bana iki veri döner string ve error mesajı string olan veriyi alıyorum erroru kontrol etmiyorum o yüzden blank identitypar kullandım
	fmt.Println("Çıktı: " + str)
}
