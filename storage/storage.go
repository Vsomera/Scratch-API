package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vsomera/scratch-api/types"
)

type MySqlStorage struct {
	db *sql.DB
}

func NewMySqlStore() (*MySqlStorage, error) {
	// attempt mysql connection
	dbPassword := os.Getenv("DB_PASSWORD")
	connString := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/fruits", dbPassword) // connection string
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MySQL ❌ : %v", err)
	}

	// ping the database to check if connected
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error pinging MySQL database ❌ : %v", err)
	}

	fmt.Println("Connected to MySQL database ✔️")
	return &MySqlStorage{db: db}, nil
}

// GetFruitByName queries the database by fruit name and returns the fruit row that matches the fruit name.
func (s *MySqlStorage) GetFruitByName(name string) (*types.Fruit, error) {
	// Write SQL query to select the fruit row based on the fruit name
	query := "SELECT id, name, count FROM fruits WHERE name = ?"

	// Execute the SQL query and retrieve the result
	row := s.db.QueryRow(query, name)

	var fruit types.Fruit
	err := row.Scan(&fruit.ID, &fruit.Name, &fruit.Count)
	if err != nil {
		return nil, err
	}

	return &fruit, nil
}

func (s *MySqlStorage) AddFruit(name string, count int) error {
	query := "INSERT INTO fruits (name, count) VALUES (?, ?)"
	_, err := s.db.Exec(query, name, count)
	if err != nil {
		return fmt.Errorf("error adding fruit: %w", err)
	}
	return nil
}
