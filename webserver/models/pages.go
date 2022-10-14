package models


type PageDefault struct {
	T string `json:"t"`
	Title string  `json:"title"`
}


type TraidingTools struct {
	Icon string `json:"icon"`
	Title string `json:"title"`
	Description string `json:"description"`
	Link string `json:"link"`
}

type Tools struct {
	Title string `json:"title"`
	Text string `json:"text"`
}

type FAQ struct {
	Question string  `json:"question"`
	Answer string `json:"answer"`
}

type PageHome struct {
	T string `json:"t"`
	Title string  `json:"title"`
	TraidingTools []TraidingTools `json:"traidingTools"`
	AllTools []Tools `json:"allTools"`
	FAQ []FAQ `json:"faq"`
}