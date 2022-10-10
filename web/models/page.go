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
	NameItem    string
}

type HomePage struct {
	ServiseBlock ServiseBlock
}


type ServiseBlock struct {
	TraidingTools []TraidingToolsItem
}

