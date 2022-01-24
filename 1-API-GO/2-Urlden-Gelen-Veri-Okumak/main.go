package main

import "net/http"

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Merhaba"))

	x := r.URL.Path[1:] // başında / olur urlden aldığım verinin başındaki / hariç alıyoyurm yani localhost:8080/ahmet i /ahmet olarak alacaktı artık ahmet olarak alır

	gelen := "hey sen " + x
	rw.Write([]byte(gelen))
}

func aboutHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("About Page"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler) //buna istek gelince aboutHandler'e gider
	http.ListenAndServe(":8080", nil)
}
