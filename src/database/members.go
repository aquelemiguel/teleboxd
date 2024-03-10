package database

func CreateMember(handle string, chatId int64) (int64, error) {
	db, _ := GetDatabase()

	// maybe the user is already in another chat
	userId, err := GetUser(handle)
	if err != nil {
		userId, _ = CreateUser(handle)
	}

	// or maybe the chat already exists
	_, err = GetChat(chatId)
	if err != nil {
		chatId, _ = CreateChat(chatId)
	}

	result, err := db.Exec("INSERT INTO members (user_id, chat_id) VALUES (?, ?)", userId, chatId)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
