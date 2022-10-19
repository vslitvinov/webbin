package modules

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"site/webserver/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)


type CartItem struct {
	Product models.Product
	Count int64
}


type Cart struct {
	Items map[string]CartItem
}

func NewCart() Cart{
	c := Cart{
		Items: make(map[string]CartItem),
	}
	return c
}


type StorageProduct interface {
	GetAllProduct(ctx context.Context) ([]models.Product, error)
}


type CartService struct {
	db StorageProduct 
	cacheProduct *Cache
	cache *Cache
}

func NewCartService(db *pgxpool.Pool) *CartService {
	cs :=  &CartService{
		db: NewDatabase(db),
		cacheProduct: NewCache(),
		cache: NewCache(),
	}

	ps,err := cs.db.GetAllProduct(context.Background())
	if err != nil {
		log.Println(err)
	}
	for _,p := range ps {
		cs.cacheProduct.Set(p.UUID,p)
	}
	log.Println(ps)


	return cs
}


// AddItemToCart добавить в корзину товар
// ResetCart очистить корзину 
// DeleteItemFromCart // удалить итем из корзины 
// GetCartInfo // получить содержимое корзины
// EditCountItem // изменить количество элемента в корзине 

type AddItemReq struct {
	UUID string `json:"uuid"`
}


type FullCart struct {
	UUID string `json:"uuid"`
	Cart     Cart
}

func (cs *CartService) AddItemToCart(w http.ResponseWriter, r *http.Request){
	// проверяем если ли токен корзинны ? 
		// если нет создаем токен и корзину
	// генерируем uuid и добавляем элемент 

	// sessionToken := uuid.NewString()
	// parse body request
	data := &AddItemReq{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		log.Println("err decode: ",err)
	}

	// read coockie file 
	var cartToken = uuid.NewString()

	token, err := r.Cookie("cart_token")
	if err != nil {
		log.Println("err get body cookie: ",err)
	} 
	if token != nil {
		cartToken = token.Value
	}

	var cart Cart

	// check cart by uuid
	c,ok := cs.cache.Get(cartToken)
	if !ok {
		cart = NewCart()
		cs.cache.Set(cartToken,cart)
	} else {
		cart = c.(Cart)
	}

	// check and get product by uuid 
	product,ok := cs.cacheProduct.Get(data.UUID)
	if !ok {
		log.Println("err not fount product by uuid")
	}


	// product item and to cart
	item, ok := cart.Items[data.UUID]
	if ok {
		item.Count += 1 
		cart.Items[data.UUID] = item
	} else {
		cart.Items[data.UUID] = CartItem{
			Product: product.(models.Product),
			Count: 1,
		}
	}

	// update cookie file 
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "cart_token",
	// 	Value:   cartToken,
	// 	Expires: time.Now().Add(1200 * time.Second),
	// })


	// 
	res, err := json.Marshal(FullCart{
		UUID: cartToken,
		Cart: cart,
	})
	if err != nil {
		log.Println("Marshal json req")
	}
	fmt.Fprintf(w,string(res))
}

func (cs *CartService) ResetCart(w http.ResponseWriter, r *http.Request){
	// удаляем содержимое корзины

	var cartToken = uuid.NewString()
	token, err := r.Cookie("cart_token")
	if err != nil {
		log.Println("err get body cookie: ",err)
	} 
	if token != nil {
		cartToken = token.Value
	}

	// clear cart 
	cs.cache.Set(cartToken, NewCart())
	
}

func (cs *CartService) DeleteItemFromCart(w http.ResponseWriter, r *http.Request){
	// фейл сейф проверка если в куках токен корзины 
	// проверка находится ли в корзине товар с нужным uuid 
		// если да удаляем 

	// parse body request
	data := &AddItemReq{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		log.Println("err decode: ",err)
	}

	// read coockie file 
	var cartToken = uuid.NewString()

	token, err := r.Cookie("cart_token")
	if err != nil {
		log.Println("err get body cookie: ",err)
	} 
	if token != nil {
		cartToken = token.Value
	}

	var cart Cart

	// check cart by uuid
	c,ok := cs.cache.Get(cartToken)
	if !ok {
		cart = NewCart()
		cs.cache.Set(cartToken,cart)
	} else {
		cart = c.(Cart)
	}

	// product item and to cart
	_, ok = cart.Items[data.UUID]
	if ok {
		delete(cart.Items,data.UUID)
	} 

	res, err := json.Marshal(FullCart{
		UUID: cartToken,
		Cart: cart,
	})
	if err != nil {
		log.Println("Marshal json req")
	}
	fmt.Fprintf(w,string(res))
}

func (cs *CartService) GetCartInfo(w http.ResponseWriter, r *http.Request){
	// read coockie file 
	var cartToken = uuid.NewString()

	token, err := r.Cookie("cart_token")
	if err != nil {
		log.Println("err get body cookie: ",err)
	} 
	if token == nil {
		res, err := json.Marshal(AddItemReq{UUID :cartToken})
		if err != nil {
			log.Println("Marshal json req")
		}
		fmt.Fprintf(w,string(res))
		return
	}
	if token != nil {
		cartToken = token.Value
	}

	var cart Cart

	// check cart by uuid
	c,ok := cs.cache.Get(cartToken)
	if !ok {
		cart = NewCart()
		cs.cache.Set(cartToken,cart)
	} else {
		cart = c.(Cart)
	}

	res, err := json.Marshal(FullCart{
		UUID: cartToken,
		Cart: cart,
	})
	if err != nil {
		log.Println("Marshal json req")
	}
	fmt.Fprintf(w,string(res))


}

func (cs *CartService) EditCountItem(w http.ResponseWriter, r *http.Request){
	// фейл сейф проверка если в куках токен корзины 
	// проверка находится ли в корзине товар с нужным uuid 
		// если да изменяем его количество 
}





