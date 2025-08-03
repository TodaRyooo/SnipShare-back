package mysql

import (
	"database/sql"
	"fmt"
	"github.com/TodaRyooo/SnipShare-back/internal/domain"
)

type commentRepositoryImpl struct {
	db *sql.DB
}

func NewCommentRepository(client *Client) domain.CommentRepository {
	return &commentRepositoryImpl{db: client.GetDB()}
}

func (r *commentRepositoryImpl) FindAll() ([]domain.Comment, error) {
	rows, err := r.db.Query("select id, snippet_id, parent_id, body, created_by, created_at FROM comment")
	if err != nil {
		return nil, fmt.Errorf("コメントの取得に失敗しました: %w", err)
	}
	defer rows.Close()

	var comments []domain.Comment
	for rows.Next() {
		var s domain.Comment
		if err := rows.Scan(&s.ID, &s.SnippetId, &s.ParentId, &s.Body, &s.CreatedBy, &s.CreatedAt); err != nil {
			fmt.Printf("行のスキャンに失敗しました。この行はスキップします: %v\n", err)
			continue
		}
		comments = append(comments, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("行のイテレーション中にエラーが発生しました: %w", err)
	}

	return comments, nil
}
