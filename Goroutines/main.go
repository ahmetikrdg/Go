package main

import "time"

//go goroutine fo tarafından çalışma zamanında yönetilen hafif bir iş parçaçığıdır.

func main() {
	//basit bir gorouting kullanımı
	go xFunc()
	time.Sleep(100 * time.Millisecond)
}

func xFunc() {
	for l := byte('a'); l <= byte('z'); l++ {
		println(string(l))
	}
}
