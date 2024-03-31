package book

import "time"

type Book struct {
	ID          string
	Name        string
	AuthorName  string
	CreaterTime time.Time
}
