package models

type Note struct {
	Title  string `json:"title"`
	Note   string `json:"note"`
	UserId int    `json:"user_id,omitempty"`
}
