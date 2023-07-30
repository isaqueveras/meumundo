package domain

import (
	"context"
	"time"
)

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Get(ctx context.Context, uf, slug *string) (*Article, error)
}

// IArticle represent the article's repository contract
type IArticle interface {
	Get(ctx context.Context, uf, slug *string) (*Article, error)
}

// Article is representing the Article data struct
type Article struct {
	ID          *string     `json:"id,omitempty"`
	CidadeID    *string     `json:"cidade_id,omitempty"`
	Conteudo    *string     `json:"conteudo,omitempty"`
	Info        interface{} `json:"info,omitempty"`
	Criacao     *time.Time  `json:"criacao,omitempty"`
	Atualizacao *time.Time  `json:"atualizacao,omitempty"`
}
