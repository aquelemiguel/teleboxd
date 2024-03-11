package database

func CreateMember(handle string, chatId int64) (int64, error) {
	db, _ := GetDatabase()

	// maybe the user is already in another chat
	userId, err := GetUser(handle)
	if err != nil {
		userId, _ = CreateUser(handle)
	}

	// or perhaps the chat already exists
	_, err = GetChat(chatId)
	if err != nil {
		chatId, _ = CreateChat(chatId)
	}

	res, err := db.Exec("INSERT INTO members (user_id, chat_id) VALUES (?, ?)", userId, chatId)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func DeleteMember(handle string, chatId int64) error {
	db, _ := GetDatabase()

	// maybe the user doesn't exist
	userId, err := GetUser(handle)
	if err != nil {
		return ErrUserNotFound
	}

	res, err := db.Exec("DELETE FROM members WHERE user_id = ? AND chat_id = ?", userId, chatId)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
