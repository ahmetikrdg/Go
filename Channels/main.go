package main

//threadler yerine go routing kullanıldığı için birde buna özel iletişim standartı gelir oda channeldir.
import "fmt"

func merhaba(chan1 chan string) { //fonksiyonumu yazdım kanalla beraber göndermek istediğim fonksiyon
	chan1 <- "Merhaba" //ok operatörü değeri kanala gönderiyo
}

func main() {
	myChannel := make(chan string) //bir channel oluşturdum. channel oşturmamın iki temel sebebi go rutinlerin birbirleriyle iletişime geçmesini sağlıyoruz
	// ikinci sebep ise herhangi bir değeri almak için return kullanamıyoduk ancak bir kanal üzerinden alabiliyoduk. çünkü return olana kadar kanaldan birşey geçmez
	//ee bir sıkıntı olursa ne olur program patlar bunun önüne geçmiş oluyoruz channel ile
	//sonra yukarıdaki fonksiyona git
	go merhaba(myChannel)    //ayrı bir go rutin oluştururken o fonksiyona parametre olarak mychannel yerleştiriyorum
	fmt.Println(<-myChannel) //kanaldeki değeri almamızı sağlıyo
}
