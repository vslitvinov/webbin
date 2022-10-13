package modules

import (
	"log"
	"net/http"
	"text/template"
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
}


func NewUsersService( ) *UsersService {

	us := &UsersService{
		cacheUser: NewCache(),
		cacheSession: NewCache(),
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
		// redirect
	}

	log.Println(r.Method)

	buf := make([]byte,500)

	n, err := r.Body.Read(buf)
	log.Println(string(buf[:n]),err)



	// err := json.NewDecoder(r.Body).Decode(&data)
	// if err != nil {
	// 	log.Println(err)
	// }

	t, err := template.ParseFiles("./webserver/template/signup.tmpl")
	if err != nil {
		log.Printf("Ошибка парсинга шаблона: %v", err)
		return
	}
	err = t.ExecuteTemplate(w,"signup.tmpl",struct{}{})
}




// Signup регистрация 
// Login 