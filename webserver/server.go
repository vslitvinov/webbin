package webserver

import (
	"html/template"
	"log"
	"net/http"
	"site/webserver/modules"
)


type WebServer struct {
	Router *http.ServeMux
	US *modules.UsersService
	Temp *template.Template
	CachePage *modules.Cache
}


func NewWebServer() *WebServer{

	var err error

	s :=  &WebServer{
		Router: http.NewServeMux(),
		US: modules.NewUsersService(),
		CachePage :  modules.NewCache(),
	}

	s.Temp, err = modules.TempLoad("./webserver/template/pages/")
	if err != nil {
		log.Println(err)
	}

	// todo struct data
	modules.LoadConfigPage(s.CachePage, "home","./webserver/config/homePage.json")

	// s.Router.HandleFunc("/signup", s.US.Signup)
	s.Router.HandleFunc("/", s.HomePage)
	s.Router.HandleFunc("/catalog", s.TradingBots)
	s.Router.HandleFunc("/signup", s.US.Signup)
	s.Router.Handle("/assets", http.FileServer(http.Dir("./web/template/")))

	return s
}


func (ws *WebServer) HomePage(w http.ResponseWriter, r *http.Request){

	data, err :=  ws.CachePage.Get("home")
	if !err {
		log.Println(err)
	}
	log.Println(data)
    ws.Temp.ExecuteTemplate(w, "Index", data)
}
func (ws *WebServer) TradingBots(w http.ResponseWriter, r *http.Request){
    ws.Temp.ExecuteTemplate(w, "Index", nil)
}