package usecase

import (
	"fmt"
	"github.com/TodaRyooo/SnipShare-back/internal/domain"
)

type CommentUsecase struct {
	repo domain.CommentRepository
}

func NewCommentUsecase(repo domain.CommentRepository) *CommentUsecase {
	return &CommentUsecase{repo: repo}
}

func (uc *CommentUsecase) GetAllComments() ([]domain.Comment, error) {
	comments, err := uc.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("コメントの取得に失敗しました: %w", err)
	}
	return comments, nil
}
