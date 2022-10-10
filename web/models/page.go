package models

type NavBarItemLink struct {
	Name string
	Link string
}

type NavBarItem struct {
	Name  string
	Links []NavBarItemLink
}

type TraidingToolsItem struct {
	Icon        string
	Title       string
	Description string
	Link    string
}

type HomePage struct {
	ServiseBlock ServiseBlock
	AllTools []Tool
	FAQ Faqs
}

type Faqs struct {
	Question string
	Answer string 
}

type Tool struct {
	Title string 
	Text string
}

type ServiseBlock struct {
	TraidingTools []TraidingToolsItem
}

