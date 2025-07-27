package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	// 各層のパッケージをインポート
	"github.com/TodaRyooo/SnipShare-back/internal/domain"
	"github.com/TodaRyooo/SnipShare-back/internal/infrastructure/mysql"
	"github.com/TodaRyooo/SnipShare-back/internal/presenter"
	"github.com/TodaRyooo/SnipShare-back/internal/usecase"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/snipshare_db?charset=utf8mb4&parseTime=true"

	mysqlClient, err := mysql.NewClient(dsn)
	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %v", err)
	}
	defer mysqlClient.Close()

	err = mysqlClient.Ping()
	if err != nil {
		log.Fatalf("データベースへのPingに失敗しました: %v", err)
	}
	fmt.Println("データベースに正常に接続しました")

	// --- 2. 各層の依存関係を注入 (DI: Dependency Injection) ---
	// まずは具体的なデータベース実装であるリポジトリを初期化
	// 次に、ビジネスロジックを担うユースケースを初期化。リポジトリに依存
	// 最後に、HTTPリクエストを処理するプレゼンター（ハンドラ）を初期化。ユースケースに依存
	var snippetRepo domain.SnippetRepository = mysql.NewSnippetRepository(mysqlClient)
	snippetUsecase := usecase.NewSnippetUsecase(snippetRepo)
	snippetPresenter := presenter.NewSnippetPresenter(snippetUsecase)

	// --- 3. ルーティングの設定 ---
	// Goの標準ルーター (http.ServeMux) を使用
	mux := http.NewServeMux()

	// 各HTTPメソッドとパスに対するハンドラ関数を登録
	// presenter層のメソッドを直接登録
	mux.HandleFunc("GET /snippets", snippetPresenter.GetSnippets) // 全スニペット取得

	// --- 4. HTTPサーバーの起動 ---
	port := ":8080"
	fmt.Printf("サーバーがポート %s で起動しました...\n", port)
	// サーバー起動。エラーが発生したらログに出してアプリケーションを終了
	log.Fatal(http.ListenAndServe(port, mux))
}
