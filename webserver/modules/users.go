package modules

import (
	"context"
	"html/template"
	"log"
	"fmt"
	"net/http"
	"site/webserver/models"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type sessionUser struct {
	email string
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
	GetAll(ctx context.Context) ([]models.User, error)
	SetUser(ctx context.Context, data models.User) (uint64, error)
}

func NewUsersService(db *pgxpool.Pool) *UsersService {

	var err error

	us := &UsersService{
		db:  NewDatabase(db), 
		cacheUser: NewCache(),
		cacheSession: NewCache(),
	}
	users,err := us.db.GetAll(context.Background())
	if err != nil {
		log.Println(err)
	}

	for _,user := range users {
		us.cacheUser.Set(user.Email,user)
	}
	log.Println("load users complite")


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

	// получаем имя пользователя из кеша сеанса
	userSession, ok := us.cacheSession.Get(c.Value)
	if !ok {
		// токен сеанса отсутствует 
		return false
	}
	// истекло ли время сессии 
	if userSession.(sessionUser).isExpired() {
		us.cacheSession.Delete(c.Value)
		return false
	}

	return true

}

// func (us *UsersService) 

func (us *UsersService) Signin (w http.ResponseWriter, r *http.Request){ 
		// если POST
	// Получение данных из формы 
	// получение данных пользователя на основе данных 
		// если такого пользователя нет возрощаем ошибку 
	// проверяем хеши паролей 
		// если нет возрат ошибки авторизации 
	// создание токена сесии и временой метки 
	// сохранение за пользователем сессии 
	// установление куков пользователю и редерект 

	if r.Method == "POST" {
		r.ParseForm()                    
    	email := r.Form.Get("email")
    	password := r.Form.Get("password")

    	user, ok := us.cacheUser.Get(email)
    	if !ok {
    		// пользователя нет
    		w.WriteHeader(http.StatusUnauthorized)
    	}
    	if !CheckPasswordHash(password, user.(models.User).Password) {
    		// пароли не совпадают
    		w.WriteHeader(http.StatusUnauthorized)
    	}

    	session := sessionUser{
    		email:  email,
    		expiry: time.Now().Add(1200 * time.Second),
    	}
    	sessionToken := uuid.NewString()
		
		us.cacheSession.Set(sessionToken,session)
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: session.expiry,
		})
		http.Redirect(w, r, "./account", http.StatusSeeOther)
	} else{
		if us.CheckSessionToken(r){
			http.Redirect(w, r, "./account", http.StatusSeeOther)
		}
	
		us.Temp.ExecuteTemplate(w, "Index", struct{
    		T string
    		Title string
    	}{
    		T: "login",
    		Title: "Login",
    	})
    }
}


func (us *UsersService) Logout(w http.ResponseWriter, r *http.Request) {
	//прекщение сессии 
}

func (us *UsersService) Account(w http.ResponseWriter, r *http.Request){
	if !us.CheckSessionToken(r){
		http.Redirect(w, r, "./login", http.StatusSeeOther)
	}

	fmt.Fprintf(w,"lol")
	// http.Redirect(w, r, newUrl, http.StatusSeeOther)
}


func (us *UsersService) Signup(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {
		u, ok := us.cacheUser.Get(r.Form.Get("email"))
		if !ok {
			// err email занят
		}
		if u.(models.User).DisplayName == r.Form.Get("displayname"){
			// err display name занят
		}
		r.ParseForm()                    
    	pass,err := HashPassword(r.Form.Get("password"))
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



func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
