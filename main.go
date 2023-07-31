package main

import (
	"github.com/go-faker/faker/v4"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func init() {
	godotenv.Load()
}

type User struct {
	Id           int    `db:"id"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	Email        string `db:"email"`
	RegisterDate string `db:"register_date"`
	InsertOrder  int    `db:"insert_order"`
}

func main() {
	db, err := sqlx.Open("postgres", os.Getenv("POSTGRES_CONNECTION"))
	panicOnError(err, "db connection")
	defer db.Close()

	var users []User
	for i := 0; i < 2000; i++ {
		users = append(users, User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			RegisterDate: faker.Date(),
			InsertOrder:  i,
		})
	}

	_, err = db.NamedExec(`insert into users(email, first_name, last_name, register_date, insert_order) 
        VALUES (:email, :first_name, :last_name, :register_date,:insert_order)`, users)

	panicOnError(err, "batch insert")

}

func panicOnError(err error, s string) {
	if err != nil {
		log.Fatalf("%s, %v", s, err)
	}

}
