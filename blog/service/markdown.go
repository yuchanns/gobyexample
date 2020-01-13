package service

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/yuchanns/gobyexample/blog/models"
	"time"
)

func autoConvertTime(m *models.Markdown) {
	m.CreatedAtTime = time.Unix(m.CreatedAt, 0).Format("2006-01-02 15:04:05")
	m.UpdatedAtTime = time.Unix(m.UpdatedAt, 0).Format("2006-01-02 15:04:05")
}

func (s *Service) Add(ctx *gin.Context) {
	var m models.Markdown

	if err := ctx.ShouldBind(&m); err != nil {
		s.JsonErr(ctx, err.Error(), 1)

		return
	}

	now := time.Now().Unix()

	m.CreatedAt = now
	m.UpdatedAt = now

	sql, args, err := sq.Insert("markdown").
		Columns("content", "is_draft", "created_at", "updated_at").
		Values(m.Content, m.IsDraft, now, now).
		ToSql()

	if err != nil {
		s.JsonErr(ctx, err.Error(), 2)

		return
	}

	_, err = s.Dao.SqlxDB().Exec(sql, args...)

	if err != nil {
		s.JsonErr(ctx, err.Error(), 3)

		return
	}

	autoConvertTime(&m)

	s.Json(ctx, &m)
}

func (s *Service) Get(ctx *gin.Context) {
	eqs := make([]sq.Eq, 2)

	id := ctx.DefaultQuery("id", "0")

	eqs[0] = sq.Eq{"is_deleted": 0}
	eqs[1] = sq.Eq{"id": id}

	q := sq.Select("*").From("markdown")

	for _, eq := range eqs {
		q.Where(eq)
	}

	query, args, err := q.ToSql()

	if err != nil {
		s.JsonErr(ctx, err.Error(), 1)

		return
	}

	var m models.Markdown

	if err := s.Dao.SqlxDB().Get(&m, query, args...); err != nil {
		s.JsonErr(ctx, err.Error(), 2)

		return
	}

	autoConvertTime(&m)

	s.Json(ctx, &m)
}
