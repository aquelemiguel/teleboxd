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

func GetUser(handle string) (User, error) {
	db, _ := GetDatabase()

	var user User
	err := db.QueryRow("SELECT * FROM users WHERE handle = ?", handle).Scan(&user.Id, &user.Handle, &user.LastLogTime)
	return user, err
}

func GetUsersByChat(chatId int64) ([]*User, error) {
	db, _ := GetDatabase()

	rows, err := db.Query("SELECT * FROM users WHERE id IN (SELECT user_id FROM members WHERE chat_id = ?)", chatId)
	if err != nil {
		return nil, err
	}

	var users []*User
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Handle, &user.LastLogTime)
		users = append(users, &user)
	}
	return users, nil
}

func GetAllUsers() ([]*User, error) {
	db, _ := GetDatabase()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []*User
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Handle, &user.LastLogTime)
		users = append(users, &user)
	}
	return users, nil
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
