package modules



type ZenOrder struct {
	TerminalUuid string 
	Amount string 
	Currency string 
	MerchantTransactionId string  //gen
	Customer ZenCustomer
	items []ZenItem
	Signature string
	ShippingAddress ZenShippingAddress
	UrlReturn string
	UrlSuccess string
	UrlFailure string
	CustomIpnUrl string
}

type ZenShippingAddress struct {
	Id string
	FirstName string
	LastName string
	Country string
	Street string
	City string
	CountryState string
	Province string
	BuildingNumber string
	RoomNumber string
	Postcode string
	CompanyName string
	Phone string
}

type ZenCustomer struct {
	FirstName string 
	LastName string 
	Email string 
}


type ZenItem struct {
	Code string
	Category string
	Name string
	Price string
	Qantity string
	LineAmountTotal string
}

type ZenPayment struct {

}

func (zp *ZenPayment) Payment() {

}