package web

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)


type Site struct {
	Router *mux.Router
} 



func NewWebSite() *Site{
	s :=  &Site{ 
		Router: mux.NewRouter(),
	}

	s.Router.HandleFunc("/signup", s.Signup)
	s.Router.HandleFunc("/login", s.Login)
	s.Router.HandleFunc("/recover-account", s.RecoverAccount)
	s.Router.HandleFunc("/account", s.Account)
	s.Router.HandleFunc("/catalog", s.TradingBots)
	s.Router.HandleFunc("/", s.HomePage)
	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/template/")))

	return s 
}

func (s *Site) Signup(w http.ResponseWriter, r *http.Request){
		t, err := template.ParseFiles("./web/template/page-signup.html")
	if err != nil {
		log.Printf("Ошибка парсинга шаблона: %v", err)
		return
	}

	err = t.Execute(w, struct{
		Hello string
		World string
	}{
		Hello: "Hello",
		World: "World",
	})
	if err != nil {
		log.Printf("Ошибка записи в шаблон: %v", err)
		return
	}
}


type HomePageItem struct {
	Icon string
	Title string
	Description string 
	NameItem string
}

func (s *Site) Login(w http.ResponseWriter, r *http.Request){}
func (s *Site) RecoverAccount(w http.ResponseWriter, r *http.Request){}
func (s *Site) Account(w http.ResponseWriter, r *http.Request){}



func (s *Site) TradingBots(w http.ResponseWriter, r *http.Request){

}



func (s *Site) HomePage(w http.ResponseWriter, r *http.Request){
		t, err := template.ParseFiles("./web/template/index.html")
	if err != nil {
		log.Printf("Ошибка парсинга шаблона: %v", err)
		return
	}



	err = t.Execute(w, struct{
		Hello string
		Items []HomePageItem
	}{
		Hello: "World",
		Items: []HomePageItem{  
			HomePageItem {
				Icon: "shared.svg",
				Title: "DCA Bot",
				Description: "No need to risk it all! Instead of investing a lump sum with unknown risks, the bot will invest the amount partially with maximum benefit.",
				NameItem: "./catalog/dca",
			},
		},
	})
	if err != nil {
		log.Printf("Ошибка записи в шаблон: %v", err)
		return
	}
}