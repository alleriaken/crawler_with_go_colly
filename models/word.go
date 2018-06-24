package models

import "fmt"

type Word struct {
	word, w_type, w_def, w_pron, examword_id  string
}

func (w *Word) NewWord(word, w_type, w_def, w_pron, examword_id string) {
	w.word = word
	w.w_def = w_def
	w.w_type = w_type
	w.w_pron = w_pron
	w.examword_id = examword_id
}

func AllWord() ([]*Word, error) {
	rows, err := db.Query("SELECT * FROM words")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words = make([]*Word, 0)

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return words, nil
}

func SaveWord(word *Word)  {
	query := "INSERT INTO words (word, w_type, w_def, w_pron, examword_id) VALUES (?, ?, ? ,?, ?) ON DUPLICATE KEY UPDATE w_def = ?"
	stmt, _ := db.Prepare(query)
	fmt.Println("Hello")
	fmt.Println(word.word)
	res, _ := stmt.Exec(word.word, word.w_type, word.w_def, word.w_pron, word.examword_id, word.w_def)
	id, _ := res.LastInsertId()
	fmt.Println(id)
}