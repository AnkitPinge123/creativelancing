package database

import (
	"billing-app/structure"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dbConn = "root:Pinge@8055@tcp(127.0.0.1:3306)/creativelancing" // TODO ankit.pinge use if from config

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("mysql", dbConn)
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}

	// Check if the database connection is valid
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Database connection established successfully")
	return nil
}

func FetchUsersFromDatabase() ([]structure.User, error) {
	rows, err := db.Query("SELECT user_id, username, password_hash, email, role, created_at, updated_at FROM Users")
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []structure.User
	for rows.Next() {
		var user structure.User
		if err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error after scanning rows: %v", err)
		return nil, err
	}

	return users, nil
}

func FetchProductsByShopID(shopID int) ([]structure.Product, error) {
	rows, err := db.Query(`
		SELECT product_id, shop_id, product_name, product_description, price, quantity_in_stock, created_at, updated_at
		FROM Products WHERE shop_id = ?`, shopID)
	if err != nil {
		log.Printf("Error querying products: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store product data
	var products []structure.Product
	for rows.Next() {
		var product structure.Product
		if err := rows.Scan(&product.ProductID, &product.ShopID, &product.ProductName, &product.ProductDesc, &product.Price, &product.QuantityInStock, &product.CreatedAt, &product.UpdatedAt); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		products = append(products, product)
	}

	// Check for errors after iterating through rows
	if err := rows.Err(); err != nil {
		log.Printf("Error after scanning rows: %v", err)
		return nil, err
	}

	return products, nil
}
