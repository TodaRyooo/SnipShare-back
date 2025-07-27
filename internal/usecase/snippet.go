package usecase

import (
	"fmt"
	"github.com/TodaRyooo/SnipShare-back/internal/domain"
)

type SnippetUsecase struct {
	repo domain.SnippetRepository
}

func NewSnippetUsecase(repo domain.SnippetRepository) *SnippetUsecase {
	return &SnippetUsecase{repo: repo}
}

func (uc *SnippetUsecase) GetAllSnippets() ([]domain.Snippet, error) {
	snippets, err := uc.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("スニペットの取得に失敗しました: %w", err)
	}
	return snippets, nil
}
