package webserver

import (
	"net/http"
	"site/webserver/modules"
	"log"
)


type WebServer struct {
	Router *http.ServeMux
	US *modules.UsersService
}


func NewWebServer() *WebServer{
	s :=  &WebServer{ 
		Router: http.NewServeMux(),
		US: modules.NewUsersService(),
	}
	
	s.Router.HandleFunc("/signup", s.US.Signup)
	s.Router.HandleFunc("/", s.HomePage)
	// s.Router.Handle("/", http.FileServer(http.Dir("./web/template/")))

	return s
}


func (ws *WebServer) HomePage(w http.ResponseWriter, r *http.Request){
 
	temp,err  := modules.TempLoad()
	if err != nil {
		log.Println(err)
	}


    s1 := temp.Lookup("page.tmpl")
    s1.ExecuteTemplate(w, "page", nil)
    // s2 := temp.Lookup("head.tmpl")
    // s2.ExecuteTemplate(w, "head", nil)
    // s3 := temp.Lookup("signup.tmpl")
    // s3.ExecuteTemplate(w, "signup", nil)

}