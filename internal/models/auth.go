package models

type SignIn struct {
	UserName string `json:"username"`
	Pass     string `json:"pass"`
}
