// internal/infrastructure/mysql/post_repository.go
package mysql

import (
	"database/sql"
	"fmt"
	"github.com/TodaRyooo/SnipShare-back/internal/domain" // ドメイン層のPostモデルとインターフェースをインポート
)

// postRepositoryImpl は domain.PostRepository インターフェースのMySQL実装です。
type postRepositoryImpl struct {
	db *sql.DB // データベース接続
}

// NewPostRepository は新しいPostRepositoryのインスタンスを生成します。
// infrastructure/mysql/client.go で定義されたClientから*sql.DBを取得して利用します。
func NewPostRepository(client *Client) domain.PostRepository {
	return &postRepositoryImpl{db: client.GetDB()}
}

// FindAll は全ての投稿をデータベースから取得し、domain.Postスライスとして返します。
func (r *postRepositoryImpl) FindAll() ([]domain.Post, error) {
	rows, err := r.db.Query("SELECT id, name, content FROM posts")
	if err != nil {
		return nil, fmt.Errorf("投稿の検索に失敗しました: %w", err)
	}
	defer rows.Close() // 必ずクローズする

	var posts []domain.Post
	for rows.Next() {
		var p domain.Post
		// データベースのカラムとGoの構造体のフィールドをマッピング
		if err := rows.Scan(&p.ID, &p.Name, &p.Content); err != nil {
			// エラーログは出すが、全てのデータが取得できなくても処理を続けるためcontinue
			fmt.Printf("行のスキャンに失敗しました。この行はスキップします: %v\n", err)
			continue
		}
		posts = append(posts, p)
	}

	// ループ中に発生したエラーをチェック
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("行のイテレーション中にエラーが発生しました: %w", err)
	}
	return posts, nil
}

// Save は新しい投稿をデータベースに保存し、挿入されたレコードのIDを返します。
func (r *postRepositoryImpl) Save(post domain.Post) (int, error) {
	// プリペアドステートメントでSQLインジェクションを防ぐ
	result, err := r.db.Exec("INSERT INTO posts (name, content) VALUES (?, ?)",
		post.Name, post.Content)
	if err != nil {
		return 0, fmt.Errorf("投稿の保存に失敗しました: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("挿入されたIDの取得に失敗しました: %w", err)
	}
	return int(id), nil
}
