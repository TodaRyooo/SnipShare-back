// internal/domain/post.go
package domain

import "database/sql" // sql.NullStringのために必要

// Post は「投稿」のドメインモデルを表します。
// これはアプリケーション内で一貫して使用されるデータ構造です。
type Post struct {
	ID      int
	Name    string
	Content sql.NullString // NULL許容カラムのため sql.NullString を使用
}

// PostRepository はPostに関するデータ永続化操作のインターフェースを定義します。
// usecase層は、このインターフェースを通じてデータ操作を行います。
// これにより、具体的なデータベース実装（例: MySQL）に依存せずに済みます。
type PostRepository interface {
	FindAll() ([]Post, error)    // 全ての投稿を取得する
	Save(post Post) (int, error) // 新しい投稿を保存し、そのIDを返す
	// FindByID(id int) (Post, error) // IDで投稿を取得する
	// Update(post Post) error      // 投稿を更新する
	// Delete(id int) error         // 投稿を削除する
}
