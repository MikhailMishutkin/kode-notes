package models

type Spell struct {
	Code int      `json:"code,omitempty"`
	Pos  int      `json:"pos,omitempty"`
	Row  int      `json:"row,omitempty"`
	Col  int      `json:"col,omitempty"`
	Len  int      `json:"len,omitempty"`
	Word string   `json:"word"`
	S    []string `json:"s,omitempty"`
}
