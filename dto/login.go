package dto

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember *bool  `json:"remember"`
}
