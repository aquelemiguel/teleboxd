package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB = nil

func initDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/teleboxd.db")
	if err != nil {
		log.Fatal("failed to open the database:", err.Error())
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			handle TEXT UNIQUE,
			last_log_time INTEGER
		)
	`)
	if err != nil {
		log.Fatal("failed to create users table:", err.Error())
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS chats (
			id INTEGER PRIMARY KEY
		)
	`)
	if err != nil {
		log.Fatal("failed to create chats table:", err.Error())
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS members (
			user_id INTEGER,
			chat_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(chat_id) REFERENCES chats(id),
			PRIMARY KEY(user_id, chat_id)
		)
	`)
	if err != nil {
		log.Fatal("failed to create members table:", err.Error())
	}

	_, err = db.Exec(`
		CREATE TRIGGER IF NOT EXISTS delete_user
		AFTER DELETE ON members
		BEGIN
			DELETE FROM users
			WHERE id = OLD.user_id AND NOT EXISTS (
				SELECT 1 FROM members
				WHERE user_id = OLD.user_id
			);
		END
	`)
	if err != nil {
		log.Fatal("failed to create delete_user trigger:", err.Error())
	}

	_, err = db.Exec(`
		CREATE TRIGGER IF NOT EXISTS delete_chat
		AFTER DELETE ON members
		BEGIN
			DELETE FROM chats
			WHERE id = OLD.chat_id AND NOT EXISTS (
				SELECT 1 FROM members
				WHERE chat_id = OLD.chat_id
			);
		END
	`)
	if err != nil {
		log.Fatal("failed to create delete_chat trigger:", err.Error())
	}

	return db, nil
}

func GetDatabase() (*sql.DB, error) {
	if DB != nil {
		return DB, nil
	}
	return initDatabase()
}
