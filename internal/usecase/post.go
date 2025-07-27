// internal/usecase/post.go
package usecase

import (
	"fmt"
	"github.com/TodaRyooo/SnipShare-back/internal/domain" // ドメイン層のPostモデルとRepositoryインターフェースをインポート
)

// PostUsecase は投稿に関するビジネスロジックをカプセル化します。
type PostUsecase struct {
	repo domain.PostRepository // domain.PostRepositoryインターフェースに依存
}

// NewPostUsecase は新しいPostUsecaseのインスタンスを生成します。
// 依存性注入により、具体的なリポジトリ実装を受け取ります。
func NewPostUsecase(repo domain.PostRepository) *PostUsecase {
	return &PostUsecase{repo: repo}
}

// GetAllPosts は全ての投稿を取得するユースケース（ビジネスロジック）です。
// 必要に応じて、ここで追加のバリデーションやデータ加工などを行います。
func (uc *PostUsecase) GetAllPosts() ([]domain.Post, error) {
	posts, err := uc.repo.FindAll() // リポジトリを通じてデータを取得
	if err != nil {
		return nil, fmt.Errorf("投稿の取得に失敗しました: %w", err)
	}
	return posts, nil
}

// CreateNewPost は新しい投稿を作成するユースケースです。
// 新しい投稿がビジネスルールに合致しているかここで検証できます。
func (uc *PostUsecase) CreateNewPost(post domain.Post) (domain.Post, error) {
	// 例: 投稿名が空でないかというビジネスルールをここでチェック
	if post.Name == "" {
		return domain.Post{}, fmt.Errorf("投稿名が空です")
	}

	id, err := uc.repo.Save(post) // リポジトリを通じてデータを保存
	if err != nil {
		return domain.Post{}, fmt.Errorf("新規投稿の作成に失敗しました: %w", err)
	}
	post.ID = id // 保存後に生成されたIDをモデルにセット
	return post, nil
}
