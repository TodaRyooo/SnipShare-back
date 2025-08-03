package domain

import (
	"time"
)

type Comment struct {
	ID        int
	SnippetId int
	ParentId  int
	Body      string
	CreatedBy int
	CreatedAt time.Time
}

type CommentRepository interface {
	FindAll() ([]Comment, error)
}
