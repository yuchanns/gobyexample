package boiler

import (
	"context"
	"database/sql"
	"github.com/bwmarrin/snowflake"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/yuchanns/gobyexample/boiler/models"
	"time"
)

var (
	Node *snowflake.Node
	Ctx  context.Context
)

const FormatIso8601 = "2006-01-02 15:04:05"

func init() {
	var err error
	Node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err.Error())
	}
	Ctx = context.Background()

	AddHookBeforeInsert()
	AddHookAfterInsert()
	AddHookAfterQuery()
	AddHookBeforeUpdate()
	AddHookAfterUpdate()
}

func AddHookBeforeInsert() {
	models.AddOrderHook(boil.BeforeInsertHook, func(ctx context.Context, exec boil.ContextExecutor, o *models.Order) error {
		t := time.Now().Unix()

		o.CreatedAt = t
		o.UpdatedAt = t

		return nil
	})
}

func AddHookAfterInsert() {
	models.AddOrderHook(boil.AfterInsertHook, func(ctx context.Context, exec boil.ContextExecutor, o *models.Order) error {
		o.CreatedAtTime = time.Unix(o.CreatedAt, 0).Format(FormatIso8601)
		o.UpdatedAtTime = time.Unix(o.UpdatedAt, 0).Format(FormatIso8601)

		return nil
	})
}

func AddHookAfterQuery() {
	models.AddOrderHook(boil.AfterSelectHook, func(ctx context.Context, exec boil.ContextExecutor, o *models.Order) error {
		o.CreatedAtTime = time.Unix(o.CreatedAt, 0).Format(FormatIso8601)
		o.UpdatedAtTime = time.Unix(o.UpdatedAt, 0).Format(FormatIso8601)

		return nil
	})
}

func AddHookBeforeUpdate() {
	models.AddOrderHook(boil.BeforeUpdateHook, func(ctx context.Context, exec boil.ContextExecutor, o *models.Order) error {
		t := time.Now().Unix()

		o.UpdatedAt = t

		return nil
	})
}

func AddHookAfterUpdate() {
	models.AddOrderHook(boil.AfterUpdateHook, func(ctx context.Context, exec boil.ContextExecutor, o *models.Order) error {
		o.CreatedAtTime = time.Unix(o.CreatedAt, 0).Format(FormatIso8601)
		o.UpdatedAtTime = time.Unix(o.UpdatedAt, 0).Format(FormatIso8601)

		return nil
	})
}

func QueryOne(db *sql.DB) (*models.Order, error) {

	order, err := models.Orders(Where("id = ?", 1)).One(Ctx, db)

	return order, err
}

func CreateOne(db *sql.DB) (*models.Order, error) {
	ctx := context.Background()
	order := models.Order{
		OrderNo:    Node.Generate().String(),
		UserID:     2,
		TotalPrice: 10000,
		Postage:    50,
		Status:     1,
		IsDeleted:  0,
	}

	err := order.Insert(ctx, db, boil.Infer())

	return &order, err
}

func UpdateOne(db *sql.DB) (int64, error) {
	order, err := QueryOne(db)

	if err != nil {
		return 0, err
	}

	rowsAff, err := order.Update(Ctx, db, boil.Infer())

	return rowsAff, err
}
