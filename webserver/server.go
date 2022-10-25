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
	CartApi *modules.CartService
	Pages map[string]models.Page
}


func NewWebServer(db *pgxpool.Pool) *WebServer{

	var err error

	s :=  &WebServer{
		Router: http.NewServeMux(),
		US: modules.NewUsersService(db),
		CartApi: modules.NewCartService(db),
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
	s.Router.HandleFunc("/grid", s.Grid)
	s.Router.HandleFunc("/hold", s.Hold)
	s.Router.HandleFunc("/signup", s.US.Signup)
	s.Router.HandleFunc("/login", s.US.Signin)
	s.Router.HandleFunc("/cart", s.Cart)
	s.Router.HandleFunc("/education", s.Education)
	s.Router.HandleFunc("/faqs", s.FAQ)
	s.Router.HandleFunc("/privacy", s.Privacy)
	s.Router.HandleFunc("/terms", s.Terms)
	s.Router.HandleFunc("/contact", s.Contact)
	s.Router.HandleFunc("/checkout", s.Checkouts)
	s.Router.HandleFunc("/cart/additem", s.CartApi.AddItemToCart)
	s.Router.HandleFunc("/cart/deleteitem", s.CartApi.DeleteItemFromCart)
	s.Router.HandleFunc("/cart/getcart", s.CartApi.GetCartInfo)


	fs := http.FileServer(http.Dir("./webserver/assets"))
    s.Router.Handle("/assets/", http.StripPrefix("/assets", fs))


	return s
}


func (ws *WebServer) HomePage(w http.ResponseWriter, r *http.Request){
	 // w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
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
func (ws *WebServer) Hold(w http.ResponseWriter, r *http.Request){
    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["hold"])
    if err != nil {
    	log.Println(err)
    }
}

func (ws *WebServer) Grid(w http.ResponseWriter, r *http.Request){
    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["grid"])
    if err != nil {
    	log.Println(err)
    }
}

func (ws *WebServer) Education(w http.ResponseWriter, r *http.Request){
    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["education"])
    if err != nil {
    	log.Println(err)
    }
}
func (ws *WebServer) FAQ(w http.ResponseWriter, r *http.Request){
    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["faqpage"])
    if err != nil {
    	log.Println(err)
    }
}

func (ws *WebServer) Privacy(w http.ResponseWriter, r *http.Request){
    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["privacy"])
    if err != nil {
    	log.Println(err)
    }
}

func (ws *WebServer) Terms(w http.ResponseWriter, r *http.Request){
    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["terms"])
    if err != nil {
    	log.Println(err)
    }
}

func (ws *WebServer) Contact(w http.ResponseWriter, r *http.Request){
    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["contact"])
    if err != nil {
    	log.Println(err)
    }
}

func (ws *WebServer) Checkouts(w http.ResponseWriter, r *http.Request){
    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["checkouts"])
    if err != nil {
    	log.Println(err)
    }
}




func (ws *WebServer) Cart(w http.ResponseWriter, r *http.Request){

    err := ws.Temp.ExecuteTemplate(w, "Index", ws.Pages["cart"])
    if err != nil {
    	log.Println(err)
    }
}











