package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vsomera/scratch-api/types"
)

type Storage interface {
	GetAllFruits() ([]*types.Fruit, error)
	AddFruit(name string, count int) error
	EditFruit(name string, count int) error
}

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

	// defer db.Close() // TODO : implement DB connections to close and open

	fmt.Println("Connected to MySQL database ✔️")
	return &MySqlStorage{db: db}, nil
}

// GetAllFruits returns all fruits from the database
func (s *MySqlStorage) GetAllFruits() ([]*types.Fruit, error) {
	// select all rows in fruits table
	query := "SELECT id, name, count FROM fruits"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// put returned rows int fruits slice
	var fruits []*types.Fruit
	for rows.Next() {
		var fruit types.Fruit
		err := rows.Scan(&fruit.ID, &fruit.Name, &fruit.Count)
		if err != nil {
			return nil, err
		}
		fruits = append(fruits, &fruit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return fruits, nil
}

// AddFruit queries database to check if a row with the same name exists, if exists it adds the count to the existing row, if not it creates a new row with the name and count
func (s *MySqlStorage) AddFruit(name string, count int) error {

	// check if a row with the same name already exists
	var existingCount int
	err := s.db.QueryRow("SELECT count FROM fruits WHERE name = ?", name).Scan(&existingCount)
	if err != nil {
		if err != sql.ErrNoRows {
			return fmt.Errorf("error checking existing fruit: %w", err)
		}
		// if no rows found, fruit doesn't exist => proceed to add a new row
	} else {
		// if a row with the same name exists => increase the count and return
		newCount := existingCount + count
		_, err := s.db.Exec("UPDATE fruits SET count = ? WHERE name = ?", newCount, name)
		if err != nil {
			return fmt.Errorf("error updating existing fruit count: %w", err)
		}
		return nil
	}

	// add a new row with the fruit and the count
	query := "INSERT INTO fruits (name, count) VALUES (?, ?)"
	_, err = s.db.Exec(query, name, count)
	if err != nil {
		return fmt.Errorf("error adding fruit: %w", err)
	}
	return nil
}

// EditFruit queries db and checks if a fruit by the same name exists, then edits the count of that fruit
func (s *MySqlStorage) EditFruit(id string, count int) error {
	// check if the fruit exists
	err := s.db.QueryRow("SELECT id FROM fruits WHERE id = ?", id).Scan(&id)

	if err != nil {
		return fmt.Errorf("error checking existing fruit: %w", err)
	} else {
		// edit the existing count in database with the new count
		_, err := s.db.Exec("UPDATE fruits SET count = ? WHERE id = ?", count, id)
		if err != nil {
			return fmt.Errorf("error editing fruit count: %s", err)
		}
		return nil
	}

}
