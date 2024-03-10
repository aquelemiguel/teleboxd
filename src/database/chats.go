package database

func CreateChat(chatId int64) (int64, error) {
	db, _ := GetDatabase()

	result, err := db.Exec("INSERT INTO chats (id) VALUES (?)", chatId)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func GetChat(chatId int64) (int64, error) {
	db, _ := GetDatabase()

	var id int64
	err := db.QueryRow("SELECT id FROM chats WHERE id = ?", chatId).Scan(&id)
	return id, err
}
