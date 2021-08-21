package articles

import (
	"encoding/json"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type new_Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

func (a Article) toString() string {
	var s string
	s = "{\"Id\":\"" + a.Id + "\",\"Title\":\"" + a.Title + "\",\"Desc\":\"" + a.Desc + "\",\"Content\":\"" + a.Content + "\"}"
	return s
}
