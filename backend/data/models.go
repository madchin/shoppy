package data

type UserDetail struct {
	Uuid      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type User struct {
	Email string `json:"email"`
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
}

type Address struct {
	Uuid       string `json:"Uuid"`
	Id         int    `json:"id"`
	PostalCode string `json:"postalCode"`
	Address    string `json:"address"`
	Country    string `json:"country"`
	City       string `json:"city"`
}

type Addresses []Address

type Phone struct {
	Uuid   string `json:"Uuid"`
	Id     int    `json:"id"`
	Number string `json:"number"`
}

type Phones []*Phone

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Price float32 `json:"price"`
}

type Products []Product

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Categories []Category

type ProductCategory struct {
	ProductId  int `json:"productId"`
	CategoryId int `json:"categoryId"`
}

type ProductCategories []ProductCategory

type Shipper struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Shippers []Shipper

type Order struct {
	Id                int `json:"id"`
	DeliveryAddressId int `json:"deliveryAddressId"`
	ProductId         int `json:"productId"`
	ShipperId         int `json:"shipperId"`
}

type Orders []Order
