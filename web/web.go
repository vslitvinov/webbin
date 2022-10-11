package web

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"site/web/models"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"encoding/json"
	"os"
)


var homePage models.HomePage 



var (
    // key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
    key = []byte("super-secret-key")
    store = sessions.NewCookieStore(key)
)



type Site struct {
	Router *mux.Router
} 



func NewWebSite() *Site{
	s :=  &Site{ 
		Router: mux.NewRouter(),
	}

	file, err := os.ReadFile("./web/data/home.json")
	if err != nil {
		log.Println(err)
	}
	
	if err := json.Unmarshal(file,&homePage); err != nil {
		log.Println(err)
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

func (s *Site) Login(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "cookie-name")

    // Аутентификация проходит здесь
    // ...
    fmt.Println(session)

    // Установить пользователя как аутентифицированного
    session.Values["authenticated"] = true
    session.Save(r, w)
}

func (s *Site) RecoverAccount(w http.ResponseWriter, r *http.Request){}
func (s *Site) Account(w http.ResponseWriter, r *http.Request){}



func (s *Site) TradingBots(w http.ResponseWriter, r *http.Request){

}



func (s *Site) HomePage(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles(
		"./web/template/head.tmpl",
		"./web/template/header.tmpl",
		"./web/template/homeBlock.tmpl",
		"./web/template/services.tmpl",
	)
	if err != nil {
		log.Printf("Ошибка парсинга шаблона: %v", err)
		return
	}


	err = t.ExecuteTemplate(w,"head.tmpl",struct{
		Title string
	}{
		Title: "Eternal Intelligence - Automated systems",
	})
	err = t.ExecuteTemplate(w,"header.tmpl",struct{}{})
	err = t.ExecuteTemplate(w,"homeBlock.tmpl",struct{}{})



	err = t.ExecuteTemplate(w,"services.tmpl", homePage.ServiseBlock)
	if err != nil {
		log.Printf("Ошибка записи в шаблон: %v", err)
		return
	}
}