package articles

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type New_Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

func (a Article) ToString() string {
	s := "{\"Id\":\"" + a.Id + "\",\"Title\":\"" + a.Title + "\",\"Desc\":\"" + a.Desc + "\",\"Content\":\"" + a.Content + "\"}"
	return s
}
