package presenter

import (
	"encoding/json"
	"log"
	"net/http"

	// ドメイン層のPostモデルをインポート
	// "github.com/TodaRyooo/SnipShare-back/internal/domain"
	"github.com/TodaRyooo/SnipShare-back/internal/usecase" // ユースケース層をインポート
)

type SnippetPresenter struct {
	usecase *usecase.SnippetUsecase
}

// NewPostPresenter は新しいPostPresenterのインスタンスを生成します。
func NewSnippetPresenter(uc *usecase.SnippetUsecase) *SnippetPresenter {
	return &SnippetPresenter{usecase: uc}
}

// NOTE: 全スニペットを返すハンドラ
func (h *SnippetPresenter) GetSnippets(w http.ResponseWriter, r *http.Request) {
	snippets, err := h.usecase.GetAllSnippets()
	if err != nil {
		log.Printf("スニペットの取得エラー: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(snippets); err != nil {
		log.Printf("JSONエンコードエラー: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
