package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Customer struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Profile struct {
	ID      int    `json:"id"`
	Cid     int    `json:"custid"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Employee struct{
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Salary  int    `json:"salary"`
	Title   string `json:"title"`
}

type Department struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Manager struct {
	Eid     int    `json:"empid"`
	Did     int    `json:"depid"`
}

type Product struct {
	ID      int    `json:"id"`
	Did     int    `json:"depid"`
	Pname   string `json:"name"`
}

type Store struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Did     int    `json:"depid"`
	Eid     int    `json:"empid"`
}

type Inventory struct {
	ID       int    `json:"id"`
	Pid      int    `json:"prodid"`
	Quantity int    `json:"quantity"`
}

type Sales struct {
	Pid	  int    `json:"prodid"`
	Sid	  int    `json:"storeid"`
	Amount  int    `json:"amount"`
	Price   int    `json:"price"`
}

type Orders struct {
	ID      int    `json:"id"`
	Cid     int    `json:"custid"`
	Date	  date   `json:"date"`
	Pid     int	   `json:"prodid"`
}


func main() {
	r := gin.Default()

	r.POST("/create-user", func(c *gin.Context) {
		var custData Customer
		if err := c.BindJSON(&custData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		db, err := sql.Open("postgres", "user=sherif password=1234 dbname=postgres sslmode=disable")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
			return
		}
		defer db.Close()

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
			return
		}

		custID, err := createCustomer(tx, custData)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		err = createProfile(tx, userID, "John Doe", "123 Main St")
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
			return
		}

		empID, err := createEmployee(tx, "James Moore", 12000, "Software engineer")
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
			return
		}

		depID, err := createDepartment(tx, "Engineering department", "123 st. 456 county")
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create department"})
			return
		}
		
		err = tx.Commit()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Transaction committed successfully"})
	})

	r.Run(":8000")
}

func createCustomer(tx *sql.Tx, customer Customer) (int, error) {
	var custID int
	err := tx.QueryRow("INSERT INTO customer (username, email) VALUES ($1, $2) RETURNING id", customer.Username, customer.Email).Scan(&custID)
	if err != nil {
		return 0, err
	}

	return custID, nil
}

func createProfile(tx *sql.Tx, custID int, name, address string) error {
	_, err := tx.Exec("INSERT INTO profiles (Cid, name, address) VALUES ($1, $2, $3)", custID, name, address)
	if err != nil {
		return err
	}


	return nil
}

func createEmployee(tx *sql.Tx, salary int, name, title string) (int, error) {
	var empID int
	err := tx.QueryRow("INSERT INTO employee (name,salary,title) VALUES ($1, $2, $3) RETURNING id", name, salary, title).Scan(&empID)
	if err != nil {
		return 0, err
	}

	return empID, nil
}

func createDepartment(tx *sql.Tx, name, address string) (int, error) {
	var depID int
	err := tx.QueryRow("INSERT INTO department (name,address) VALUES ($1, $2) RETURNING id", name, address).Scan(&depID)
	if err != nil {
		return 0, err
	}

	return depID, nil
}

