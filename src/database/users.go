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

	res, err := db.Exec("INSERT INTO users (handle, last_log_time) VALUES (?, ?)", handle, now)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func GetUser(handle string) (int64, error) {
	db, _ := GetDatabase()

	var id int64
	err := db.QueryRow("SELECT id FROM users WHERE handle = ?", handle).Scan(&id)
	return id, err
}

func UpdateUser(handle string, lastLogTime int64) (int64, error) {
	db, _ := GetDatabase()

	res, err := db.Exec("UPDATE users SET last_log_time = ? WHERE handle = ?", lastLogTime, handle)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func DeleteUser(handle string) (int64, error) {
	db, _ := GetDatabase()

	res, err := db.Exec("DELETE FROM users WHERE handle = ?", handle)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}
