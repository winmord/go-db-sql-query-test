package awesomeProject5

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

type Client struct {
	ID       int
	FIO      string
	Login    string
	Birthday string
	Email    string
}

func selectClient(db *sql.DB, id int) (Client, error) {
	cl := Client{}

	err := db.QueryRow("SELECT id, fio, login, birthday, email FROM clients WHERE id = ?", id).
		Scan(&cl.ID, &cl.FIO, &cl.Login, &cl.Birthday, &cl.Email)
	if err != nil {
		return cl, err
	}

	return cl, nil
}

func insertClient(db *sql.DB, client Client) (int, error) {
	res, err := db.Exec("INSERT INTO clients (fio, login, birthday, email) VALUES (?, ?, ?, ?)",
		client.FIO, client.Login, client.Birthday, client.Email)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func deleteClient(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM clients WHERE id = ?", id)

	return err
}

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД

	clientID := 1

	// напиши тест здесь
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД

	clientID := -1

	// напиши тест здесь
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
}
