package domain

import (
	"context"
	"github.com/yuchanns/gobyexample/ddd/app/dto"
)

type NewsDomain struct {
	newsRepo INewsRepo
}

func NewNewsDomain(repo INewsRepo) *NewsDomain {
	return &NewsDomain{
		newsRepo: repo,
	}
}

func (d *NewsDomain) Add(ctx context.Context, dto *dto.NewsDTO) error {
	do := NewNews(dto.ID, dto.Title, dto.Content)
	return d.newsRepo.Add(ctx, do)
}
