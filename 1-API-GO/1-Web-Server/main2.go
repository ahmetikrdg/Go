package main

import "net/http"

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Merhaba"))
	rw.WriteHeader(http.StatusOK) //istemciden sunucuya gönderilen isteğin sunucu tarafında başarılı şekilde yapıldığını belirtir
}

func aboutHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("About Page"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler) //buna istek gelince aboutHandler'e gider
	http.ListenAndServe(":8080", nil)
}
