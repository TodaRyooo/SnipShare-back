package mysql

import (
	"database/sql"
	"fmt"
	"github.com/TodaRyooo/SnipShare-back/internal/domain"
)

type snippetRepositoryImpl struct {
	db *sql.DB
}

func NewSnippetRepository(client *Client) domain.SnippetRepository {
	return &snippetRepositoryImpl{db: client.GetDB()}
}

func (r *snippetRepositoryImpl) FindAll() ([]domain.Snippet, error) {
	rows, err := r.db.Query("SELECT id, name, body, type_id, created_at, updated_at, created_by FROM snippet")
	if err != nil {
		return nil, fmt.Errorf("投稿の検索に失敗しました: %w", err)
	}
	defer rows.Close()

	var snippets []domain.Snippet
	for rows.Next() {
		var s domain.Snippet
		if err := rows.Scan(&s.ID, &s.Name, &s.Body, &s.TypeId, &s.CreatedAt, &s.UpdateAt, &s.CreatedBy); err != nil {
			fmt.Printf("行のスキャンに失敗しました。この行はスキップします: %v\n", err)
			continue
		}
		snippets = append(snippets, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("行のイテレーション中にエラーが発生しました: %w", err)
	}
	return snippets, nil
}
