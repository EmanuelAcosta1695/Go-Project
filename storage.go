package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error // return an error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)     // return a slice of pointers to Account and an error
	GetAccountByID(int) (*Account, error) // return a pointer to Account and an error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil // nil is the error
}

func (s *PostgresStorage) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStorage) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key, 
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance numeric,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateAccount(account *Account) error {
	query := `INSERT INTO account (first_name, last_name, number, balance, created_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`

	// Execute the query and scan the generated ID into the account
	err := s.db.QueryRow(
		query,
		account.FirstName,
		account.LastName,
		account.Number,
		account.Balance,
		account.CreatedAt,
	).Scan(&account.ID)

	if err != nil {
		fmt.Printf("Error inserting account: %v\n", err)
		return err
	}

	// Print the account data for debugging purposes
	fmt.Printf("Created account: %+v\n", account)

	return nil
}

func (s *PostgresStorage) UpdateAccount(acount *Account) error {
	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStorage) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStorage) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
