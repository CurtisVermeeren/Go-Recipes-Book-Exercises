package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Product provides a data model for the productstore
type Product struct {
	ID          int
	Title       string
	Description string
	Price       float32
}

// db provides a connection to the database
var db *sql.DB

// init creates the db
func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Open db
	db, err = sql.Open("postgres", os.Getenv("POSTGRES_CONNECTION"))
	if err != nil {
		log.Fatalf("database can't connect: %v", err)
	}
	// Ping db
	err = db.Ping()
	if err != nil {
		log.Fatalf("could not ping database: %v", err)
	}
}

func main() {
	product := Product{
		Title:       "Amazon Echo",
		Description: "Amazon Echo - Black",
		Price:       179.99,
	}
	product2 := Product{
		Title:       "Staples Hyken",
		Description: "Mesh Chair - Black",
		Price:       289.99,
	}
	// insert the product
	createProduct(product)
	createProduct(product2)
	// read the products
	getProducts()
}

// createProduct inserts prd into the database and prints details of the insertion
func createProduct(prd Product) {
	result, err := db.Exec("INSERT INTO products(title, description, price) VALUES($1, $2, $3)", prd.Title, prd.Description, prd.Price)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Created successfully (%d row affected)\n", rowsAffected)
}

//
func getProducts() {
	// Get all products and check for errors
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		// No rows were found
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close()
	// Get all products and store them
	var products []*Product
	for rows.Next() {
		var id int
		prd := &Product{}
		err := rows.Scan(&id, &prd.Title, &prd.Description, &prd.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, prd)
	}
	// Check for errors
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	// Print details for all products
	for _, pr := range products {
		fmt.Printf("%s, %s, $%.2f\n", pr.Title, pr.Description, pr.Price)
	}
}
