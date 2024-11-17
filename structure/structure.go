package structure

type User struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password_hash"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Product struct {
	ProductID       int     `json:"product_id"`
	ShopID          int     `json:"shop_id"`
	ProductName     string  `json:"product_name"`
	ProductDesc     string  `json:"product_description"`
	Price           float64 `json:"price"`
	QuantityInStock int     `json:"quantity_in_stock"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}
