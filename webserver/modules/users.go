package modules

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"site/webserver/models"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
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
	db Storage
	cacheUser *Cache
	cacheSession *Cache
	Temp *template.Template
}


type Storage interface {
	Get(ctx context.Context) ([]models.User, error)
	SetUser(ctx context.Context, data models.User) (uint64, error)
}

func NewUsersService(db *pgxpool.Pool) *UsersService {

	var err error

	us := &UsersService{
		db:  NewDatabase(db), 
		cacheUser: NewCache(),
		cacheSession: NewCache(),
	}

	us.Temp, err = TempLoad("./webserver/template/auth/")
	if err != nil {
		log.Println(err)
	}

	return us

} 

func (us *UsersService) GetUserByUsername(username string){
	// us.cacheUser.Get()
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

func (us *UsersService) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func (us *UsersService) CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (us *UsersService) Login(w http.ResponseWriter, r *http.Request) {
	if us.CheckSessionToken(r) {
		 http.Redirect(w, r, "./account", http.StatusSeeOther)
	}

	if r.Method == "POST" {

		r.ParseForm()                    
    	email := r.Form.Get("email")
    	password := r.Form.Get("password")
    	log.Println(email,password)
	}


    us.Temp.ExecuteTemplate(w, "Index", struct{
    	T string
    	Title string
    }{
    	T: "login",
    	Title: "Login",
    })


	expiresAt := time.Now().Add(1000 * time.Second)
	http.SetCookie(w, &http.Cookie{
		Name:    "Test",
		Value:   "test12",
		Expires: expiresAt,
	})
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

	if r.Method == "POST" {
		// check username and email 
		// if 
		r.ParseForm()                    
    	pass,err := us.HashPassword(r.Form.Get("password"))
    	if err != nil {
    		log.Println(err)
    	}
    	user := models.User{
    		FirstName: r.Form.Get("firstname"),
			LastName:  r.Form.Get("lastname"),
			DisplayName: r.Form.Get("displayname"),
			Email: r.Form.Get("email"),
			Password: pass,
    	}

    	user_id, err := us.db.SetUser(context.Background(),user)
    	if err != nil {
    		log.Println(err)
    	}

    	user.ID = user_id

    	us.cacheUser.Set(user.DisplayName,user)

    	http.Redirect(w, r, "./account", http.StatusSeeOther)

	}






    us.Temp.ExecuteTemplate(w, "Index", struct{
    	T string
    	Title string
    }{
    	T: "signup",
    	Title: "Signup",
    })


}




// Signup регистрация 
// Login 