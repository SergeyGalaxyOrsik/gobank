package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgreStore struct {
	db *sql.DB
}

func NewPostgreStore() (*PostgreStore, error) {
	connStr := "user=sglx dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreStore{
		db: db,
	}, nil
}

func (s *PostgreStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgreStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		firstName varchar(50),
		lastName varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgreStore) CreateAccount(acc *Account) error {
	query := `insert into account
	(firstName, lastName, number, balance, created_at)
	values ($1, $2, $3, $4, $5)`

	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt)
	if err != nil {
		return nil
	}

	fmt.Printf("%+v\n", resp)
	return nil
}

func (s *PostgreStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgreStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgreStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
