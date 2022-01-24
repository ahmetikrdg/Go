package dataloaders

import (
	"encoding/json"

	. "../models" //noktanın anlamı models içindeki tüm klasörleri buraya yükle ben bunları kullanıcam .. ise ana dizin
	util "../utils"
)

func LoadUsers() []User { //kullanıcıları kullanıclar dosyasından userjson'dan alacak ve usernesnesi tipine dönüştürecek json formatındaki veriyi ve bu şekilde geriye dönecek
	bytes, _ := util.ReadFile("../json/users.json")
	var data []User
	json.Unmarshal([]byte(bytes), &data)
	return data
}

func LoadInterest() []Interest {
	bytes, _ := util.ReadFile("../json/interest.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}

func LoadInterestMapping() []InterestMapping {
	bytes, _ := util.ReadFile("../json/UserInterestMappings.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
