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

func GetChatsByUser(handle string) ([]int64, error) {
	db, _ := GetDatabase()

	rows, err := db.Query("SELECT chat_id FROM members WHERE user_id = (SELECT id FROM users WHERE handle = ?)", handle)
	if err != nil {
		return nil, err
	}

	var chats []int64
	for rows.Next() {
		var chatId int64
		rows.Scan(&chatId)
		chats = append(chats, chatId)
	}
	return chats, nil
}
