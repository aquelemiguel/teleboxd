package database

import (
	"time"
)

type User struct {
	Id          int64
	Handle      string
	LastLogTime int64
}

func CreateUser(handle string) (int64, error) {
	db, _ := GetDatabase()
	now := time.Now().Unix()

	result, err := db.Exec("INSERT INTO users (handle, last_log_time) VALUES (?, ?)", handle, now)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func GetUser(handle string) (int64, error) {
	db, _ := GetDatabase()

	var id int64
	err := db.QueryRow("SELECT id FROM users WHERE handle = ?", handle).Scan(&id)
	return id, err
}

func DeleteUser(handle string) error {
	db, _ := GetDatabase()

	_, err := db.Exec("DELETE FROM users WHERE handle = ?", handle)
	if err != nil {
		return err
	}

	return nil
}
