package modules

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type sessionUser struct {
	username string
	expiry   time.Time
}

// определяем, истек ли срок действия сеанса.
func (s sessionUser) isExpired() bool {
	return s.expiry.Before(time.Now())
}

type UsersService struct {
	cacheUser *Cache
	cacheSession *Cache
	Temp *template.Template
}


func NewUsersService( ) *UsersService {

	var err error

	us := &UsersService{
		cacheUser: NewCache(),
		cacheSession: NewCache(),
	}

	us.Temp, err = TempLoad("./webserver/template/auth/")
	if err != nil {
		log.Println(err)
	}

	return us

} 

func (us *UsersService) CheckSessionToken(r *http.Request) bool {
	// токен сеанса из cookies
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return false
		}
		return false
	}

	sessionToken := c.Value

	// получаем имя пользователя из кеша сеанса
	userSession, ok := us.cacheSession.Get(sessionToken)
	if !ok {
		// токен сеанса отсутствует 
		return false
	}
	// истекло ли время сессии 
	if userSession.(sessionUser).isExpired() {
		us.cacheSession.Delete(sessionToken)
		return false
	}

	return true

}

func (us *UsersService) Logout(w http.ResponseWriter, r *http.Request) {
	//прекщение сессии 
}

func (us *UsersService) Account(w http.ResponseWriter, r *http.Request){
	// http.Redirect(w, r, newUrl, http.StatusSeeOther)
}


func (us *UsersService) Signup(w http.ResponseWriter, r *http.Request){
	if us.CheckSessionToken(r) {
		// http.Redirect(w, r, newUrl, http.StatusSeeOther)
	}

	r.ParseForm()                    
    x := r.Form.Get("lastname")


	log.Println(x)
	log.Println(r.Method)

	buf := make([]byte,500)

	n, err := r.Body.Read(buf)
	log.Println(string(buf[:n]),err)


	if r.Method == "POST" {
		log.Println(r.Body)
	}


    us.Temp.ExecuteTemplate(w, "Index", struct{
    	Title string
    }{
    	Title: "Signup",
    })


}




// Signup регистрация 
// Login 