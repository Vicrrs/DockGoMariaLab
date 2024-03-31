package storage

import "github.com/Vicrrs/first_APIREST_go/internal/book"

type Storage interface {
	Create(book.Book) error
}
