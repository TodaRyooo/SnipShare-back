package domain

import (
	// "database/sql" // sql.NullStringのために必要
	"time"
)

type Snippet struct {
	ID        int
	Name      string
	Body      string
	TypeId    int
	CreatedAt time.Time
	UpdateAt  time.Time
	CreatedBy int
}

type SnippetRepository interface {
	FindAll() ([]Snippet, error)
	// Save(snippet Snippet) (int, error) // 新しい投稿を保存し、そのIDを返す
}
