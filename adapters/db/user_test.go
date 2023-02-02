package db_test

import (
	"database/sql"
	"log"
	"testing"
	"wwchacalww/go-cem304/adapters/db"
	"wwchacalww/go-cem304/application"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var DB *sql.DB

func setUp() {
	DB, _ = sql.Open("sqlite3", ":memory:")

	createTable(DB)
	createUsers(DB)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE users (
		"id" string,
		"name" string,
		"email" string,
		"password" string,
		"role" string,
		"status" boolean
	);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createUsers(db *sql.DB) {
	insert := `insert into users values
	("id-1", "Fulano", "fulano@gmail.com", "password", "role-test", 1),
	("id-2", "Siclano", "siclano@gmail.com", "password", "role-test", 1);`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestUserDB_Get(t *testing.T) {
	setUp()
	defer DB.Close()
	userDB := db.NewUserDB(DB)
	user, err := userDB.FindById("id-1")
	require.Nil(t, err)
	require.Equal(t, "Fulano", user.GetName())
}

func TestUserDB_GetEmail(t *testing.T) {
	setUp()
	defer DB.Close()
	userDB := db.NewUserDB(DB)
	user, err := userDB.FindByEmail("fulano@gmail.com")
	require.Nil(t, err)
	require.Equal(t, "Fulano", user.GetName())
}

func TestUserDB_List(t *testing.T) {
	setUp()
	defer DB.Close()
	userDB := db.NewUserDB(DB)
	users, err := userDB.List()
	require.Nil(t, err)
	require.Equal(t, len(users), 2)
}

func TestUserDB_Create(t *testing.T) {
	setUp()
	defer DB.Close()
	userDB := db.NewUserDB(DB)

	user := application.NewUser()
	user.Name = "Beltrano"
	user.Email = "beltrano@mail.com"
	user.Password = "12345"
	user.Role = "atendente"
	user.Status = true

	err := userDB.Create(user)
	require.Nil(t, err)
}
