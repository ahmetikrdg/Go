package main

import (
	"fmt"
	"sync"
	"time"
)

func fonksiyon1(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Fonk1 çalıştı")

	wg.Done()
}

func fonksiyon2(wg *sync.WaitGroup) { //* pointer veri tipiyle gönderileni yakaladık
	time.Sleep(3 * time.Second)
	fmt.Println("Fonk2 çalıştı")

	wg.Done()
}

func main() {
	var vg sync.WaitGroup
	vg.Add(2) //işlem sayısını belirttim iki dedim
	//waitgroup kullanmamız için go kullanmamız erek goroutine için

	go fonksiyon2(&vg) //& ile pointerı parametre olarak adresinide gönderebiliyorum. bu sayede pass by value değil pass by referance olarak işlem görüyoruz
	go fonksiyon1(&vg)

	fmt.Println("Merhaba")

	vg.Wait()

	fmt.Println("WaitGrouplar tamamlandı")
} //ihtiyaç duyduğumuz noktalarda waitgroup kullanarak bekletme işlemi yapabiliriz
