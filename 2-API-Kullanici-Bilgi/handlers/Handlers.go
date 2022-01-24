package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "../dataloaders"
	. "../models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//page nesnesi
	page := Page{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}
	//dataloaders
	users := LoadUsers()
	interests := LoadInterest()
	interestsMappings := LoadInterestMapping()
	//işlem
	var newUsers []User
	for _, user := range users {
		for _, InterestMapping := range interestsMappings {
			if user.ID == InterestMapping.UserID {
				for _, interest := range interests {
					if InterestMapping.InterestID == interest.ID {
						user.Interest = append(user.Interest, interest)
					}
				}
			}
		}

		newUsers = append(newUsers, user)
		fmt.Println(user.FirstName)
	}

	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))
}
