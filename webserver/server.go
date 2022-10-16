package webserver

import (
	"html/template"
	"log"
	"net/http"
	"site/webserver/models"
	"site/webserver/modules"

	"github.com/jackc/pgx/v4/pgxpool"
)


type WebServer struct {
	Router *http.ServeMux
	US *modules.UsersService
	Temp *template.Template
	Pages map[string]models.Page
}


func NewWebServer(db *pgxpool.Pool) *WebServer{

	var err error

	s :=  &WebServer{
		Router: http.NewServeMux(),
		US: modules.NewUsersService(db),
	}

	s.Temp, err = modules.TempLoad("./webserver/template/pages/")
	if err != nil {
		log.Println(err)
	}

	// todo
	s.Pages,err  = modules.LoadConfigFile("./webserver/config/pages.json")

	if err != nil {
		log.Println(err)
	}

	s.Router.HandleFunc("/", s.HomePage)
	// s.Router.HandleFunc("/signup", s.US.Signup)
	s.Router.HandleFunc("/dca", s.DCA)
	s.Router.HandleFunc("/signup", s.US.Signup)
	s.Router.HandleFunc("/login", s.US.Login)


	fs := http.FileServer(http.Dir("./webserver/assets"))
    s.Router.Handle("/assets/", http.StripPrefix("/assets", fs))


	return s
}


func (ws *WebServer) HomePage(w http.ResponseWriter, r *http.Request){

    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["home"])
    if err != nil {
    	log.Println(err)
    }
}
func (ws *WebServer) DCA(w http.ResponseWriter, r *http.Request){
    ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["dca"])
}











