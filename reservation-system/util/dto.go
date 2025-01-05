package util

type User struct {
	Id       string `json:id`
	Name     string `json:name`
	MobileNo string `json:mobileNo`
	StoreId  string `json:storeId`
}

type Store struct {
	Id    string `json:id`
	Name  string `json:name`
	Limit int    `json:limit`
	Queue chan string
}

type UserResponse struct {
	UserId  string `json:"userId"`
	Name    string `json:"name"`
	MobileNo string `json:"mobileNo"`
	Message string `json:"message"`
}