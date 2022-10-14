package webserver

import (
	"html/template"
	"log"
	"net/http"
	"site/web/models"
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

	// todo
	modules.LoadConfigPage(s.CachePage, "home","./webserver/config/homePage.json",&models.HomePage{})
	s.Router.HandleFunc("/", s.HomePage)
	// s.Router.HandleFunc("/signup", s.US.Signup)
	s.Router.HandleFunc("/catalog", s.TradingBots)
	s.Router.HandleFunc("/signup", s.US.Signup)


	fs := http.FileServer(http.Dir("./webserver/assets"))
    s.Router.Handle("/assets/", http.StripPrefix("/assets", fs))


	return s
}


func (ws *WebServer) HomePage(w http.ResponseWriter, r *http.Request){

	data, err :=  ws.CachePage.Get("home")
	if !err {
		log.Println(err)
	}
    ws.Temp.ExecuteTemplate(w, "Index", data)
}
func (ws *WebServer) TradingBots(w http.ResponseWriter, r *http.Request){
    ws.Temp.ExecuteTemplate(w, "Index", nil)
}