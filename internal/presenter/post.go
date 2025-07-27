// internal/presenter/post.go
package presenter

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TodaRyooo/SnipShare-back/internal/domain"  // ドメイン層のPostモデルをインポート
	"github.com/TodaRyooo/SnipShare-back/internal/usecase" // ユースケース層をインポート
)

// PostPresenter はHTTPリクエストを処理し、JSONレスポンスを生成する構造体です。
type PostPresenter struct {
	usecase *usecase.PostUsecase // usecase.PostUsecaseに依存
}

// NewPostPresenter は新しいPostPresenterのインスタンスを生成します。
func NewPostPresenter(uc *usecase.PostUsecase) *PostPresenter {
	return &PostPresenter{usecase: uc}
}

// GetPosts は全ての投稿を取得し、JSON形式で返すHTTPハンドラです。
func (h *PostPresenter) GetPosts(w http.ResponseWriter, r *http.Request) {
	// ユースケース層のビジネスロジックを呼び出す
	posts, err := h.usecase.GetAllPosts()
	if err != nil {
		log.Printf("投稿の取得エラー: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// レスポンスヘッダの設定
	w.Header().Set("Content-Type", "application/json")
	// 取得した投稿データをJSONにエンコードしてレスポンスボディに書き込む
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		log.Printf("JSONエンコードエラー: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// CreatePost は新しい投稿を作成し、結果をJSON形式で返すHTTPハンドラです。
func (h *PostPresenter) CreatePost(w http.ResponseWriter, r *http.Request) {
	var reqPost domain.Post // リクエストボディから受け取るデータのための構造体

	// リクエストボディのJSONデータをGoの構造体にデコード
	if err := json.NewDecoder(r.Body).Decode(&reqPost); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// ユースケース層のビジネスロジックを呼び出し、新しい投稿を作成
	createdPost, err := h.usecase.CreateNewPost(reqPost)
	if err != nil {
		log.Printf("投稿の作成エラー: %v", err)
		// ユースケースからのエラーメッセージをそのままクライアントに返すことも可能（セキュリティに注意）
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスヘッダとHTTPステータスコードの設定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created はリソース作成成功を示す

	// 作成された投稿データをJSONにエンコードしてレスポンスボディに書き込む
	if err := json.NewEncoder(w).Encode(createdPost); err != nil {
		log.Printf("JSONエンコードエラー: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
