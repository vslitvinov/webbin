package models 


type Product struct {
	UUID string  `json:"uuid"`
	Name string  `json:"name"`
	Price int `json:"price"`
	Сategory string `json:"category"`
	Subscription bool `json:"subscription"`
	Level int `json:"level"`
	Icon string `json:"icon"`
	Url string `json:"url"`
}