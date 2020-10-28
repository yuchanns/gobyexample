package domain

import "context"

type INewsRepo interface {
	Add(ctx context.Context, do *News) error
	GetByID(ctx context.Context, ID int) (*News, error)
	Update(ctx context.Context, do *News) error
	RemoveByID(ctx context.Context, ID int) error
}

type NewsRepo struct{}

func (r NewsRepo) Add(ctx context.Context, do *News) error {
	// todo: add into database
	return nil
}

func (r NewsRepo) GetByID(ctx context.Context, ID int) (*News, error) {
	// todo: get from database
	return nil, nil
}

func (r NewsRepo) Update(ctx context.Context, do *News) error {
	// todo: update into database
	return nil
}

func (r NewsRepo) RemoveByID(ctx context.Context, ID int) error {
	// todo: remove from database
	return nil
}
