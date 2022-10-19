package models

type Price struct {
	Title    string      `json:"title"`
	SubTitle string      `json:"subtitle"`
	Items    []PriceItem `json:"items"`
}

type PriceItem struct {
	UUID     string   `json:"uuid"`
	Name     string   `json:"name"`
	SubTitle string   `json:"subtitle"`
	Profit   int      `json:"profit"`
	Price    int      `json:"price"`
	CountBot int      `json:"countbot"`
	TyBot    string   `json:"tybot"`
	Button   string   `json:"button"`
	Items    []string `json:"items"`
}

type FAQ struct {
	Items []FAQItem `json:"items"`
}

type FAQItem struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Tools struct {
	Title    string      `json:"title"`
	SubTitle string      `json:"subtitle"`
	Items    []ToolsItem `json:"items"`
}

type ToolsItem struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type BotHome struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Image    string `json:"image"`
}

type Need struct {
	Title    string      `json:"title"`
	SubTitle string      `json:"subtitle"`
	Items    []NeedItems `json:"items"`
}

type NeedItems struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type AdvantageLine struct {
	Title    string   `json:"title"`
	SubTitle string   `json:"subtitle"`
	Items    []string `json:"items"`
}

type BotToolsItems struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BotTools struct {
	Items []BotToolsItems `json:"items"`
}

type BotDescription struct {
	Title string   `json:"title"`
	Items []string `json:"items"`
}

type Page struct {
	NameTemp       string         `json:"nametemp"`
	Title          string         `json:"title"`
	Tools          Tools          `json:"tools"`
	Need           Need           `json:"need"`
	FAQ            FAQ            `json:"faq"`
	Price          Price          `json:"price"`
	BotHome        BotHome        `json:"botHome"`
	BotTools       BotTools       `json:"botTools"`
	BotDescription BotDescription `json:"botdescription"`
	AdvantageLine  AdvantageLine  `json:"advantageline"`
}
