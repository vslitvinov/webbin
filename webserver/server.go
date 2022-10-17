package webserver

import (
	"encoding/json"
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
	s.Router.HandleFunc("/login", s.US.Signin)
	s.Router.HandleFunc("/cart", s.Cart)


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
    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["dca"])
    if err != nil {
    	log.Println(err)
    }
}


type Rwe struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (ws *WebServer) Cart(w http.ResponseWriter, r *http.Request){



	data := &Rwe{}
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(data)
		if err != nil {
			log.Println(err)
		}
		log.Println(data)
		// var items []models.AddCartItem{}
		// json.Unmarshal(body, &items)

	}

    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["cart"])
    if err != nil {
    	log.Println(err)
    }
}











