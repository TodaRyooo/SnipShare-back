// internal/infrastructure/mysql/client.go
package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // MySQLドライバのインポート（利用のみのため_）
)

// Client はデータベース接続を管理する構造体です。
type Client struct {
	db *sql.DB
}

// NewClient は新しいMySQLデータベースクライアントを初期化して返します。
// DSN（Data Source Name）は、データベースへの接続情報を含みます。
func NewClient(dsn string) (*Client, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("データベース接続のオープンに失敗しました: %w", err)
	}
	// ここではPingは行いません。接続プールを作成するだけです。
	return &Client{db: db}, nil
}

// Ping はデータベースへの接続が有効かテストします。
func (c *Client) Ping() error {
	return c.db.Ping()
}

// Close はデータベース接続をクローズします。
func (c *Client) Close() error {
	return c.db.Close()
}

// GetDB は内部の *sql.DB オブジェクトを公開します。
// これを使って、具体的なリポジトリ実装がデータベース操作を行います。
func (c *Client) GetDB() *sql.DB {
	return c.db
}
