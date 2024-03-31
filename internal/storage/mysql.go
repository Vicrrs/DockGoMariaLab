package storage

import "github.com/Vicrrs/first_APIREST_go/internal/book"

type MysqlStorage struct {
}

func (s MysqlStorage) Create(b book.Book) error {
	return nil
}
