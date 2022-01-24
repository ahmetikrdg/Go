package utils

func IsEmpty(data string) bool { //filename olarak gelen stringi kontrol edeceğim eğer boş değilse kullanılabilir hala getireceğim boşsa kullanılmaycak
	if len(data) == 0 {
		return true
	}
	return false

}
