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
	cache *Cache
}


func NewUsersService( ) UsersService {

	us := UsersService{
		cache: NewCache(),
	}

	return us

}


