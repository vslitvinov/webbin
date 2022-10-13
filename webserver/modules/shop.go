package modules

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)





type ShopService struct {
	cacheCart *Cache
	cacheOrder *Cache
}


func NewShopService () *ShopService {
	s := &ShopService{
		cacheCart: NewCache(),
		cacheOrder: NewCache(),
	}
	return s
}


func (s *ShopService) AddItemToCart(w http.ResponseWriter, r *http.Request){

	// есть ли токен корзины 

	// var cart Cart
	var cartToken string

	c, err := r.Cookie("cart_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// создаем токен 
			cartToken = uuid.NewString()
			s.cacheCart.Set(cartToken, NewCart())
		}
		log.Println("Error read cookie")
	} else {
		cartToken = c.Value
	}

	// data,ok := s.cacheCart.Get(cartToken)
	// if !ok {
	// 	log.Println("Error get cart by token")
	// }

	// cart = data.(Cart)



	// uuidItem := cart.Set()

	// нет создаем корзину 

	//  добавлем товар 




}




//Функция GetCartInfo вернет весь товар в корзине с общей стоимостью
//Функция AddItemToCart добавит товар в корзину и вернет вставленный идентификатор
//Функция «Удалить товар из корзины» удалит товар из корзины и вернет номер удаленного товара с идентификатором удаленного товара
// Функция ResetCart удалит все товары из корзины и вернет количество удаленных товаров.