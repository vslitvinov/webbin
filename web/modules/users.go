package modules



type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type User struct {
	FirstName string
	LastName  string
	DisplayName string
	Email string
	Password string
}


type UsersService struct {
	// db Storage
	cacheUser *Cache
	cacheSession *Cache
}


func NewUsersService( ) UsersService {

	us := UsersService{
		cacheUser: NewCache(),
		cacheSession: NewCache(),
	}

	return us

} 

// Signup регистрация 
// Login 


func (us *UsersService) CheckPassword(username string, pass string) bool{

	user,ok := us.cache.Get(username)
	if !ok {
		return false
	}
	// todo hash password 
	if user.Password == pass 

}