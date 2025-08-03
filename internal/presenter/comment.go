package presenter

import (
	"encoding/json"
	"github.com/TodaRyooo/SnipShare-back/internal/usecase" // ユースケース層をインポート
	"log"
	"net/http"
)

type CommentPresenter struct {
	usecase *usecase.CommentUsecase
}

func NewCommentPresenter(uc *usecase.CommentUsecase) *CommentPresenter {
	return &CommentPresenter{usecase: uc}
}

func (h *CommentPresenter) GetComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.usecase.GetAllComments()
	if err != nil {
		log.Printf("コメントの取得エラー: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		log.Printf("JSONエンコードエラー: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
