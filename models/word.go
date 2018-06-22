package models

type Word struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
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