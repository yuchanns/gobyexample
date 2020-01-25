package squirrel

import (
	sq "github.com/Masterminds/squirrel"
)

func Get(dao *Dao, id int) (*Order, error) {
	var order Order

	sql, args, err := sq.Select("*").
		From(order.TableName()).
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return &order, err
	}

	return &order, dao.DB().Get(&order, sql, args...)
}

func Insert(dao *Dao, order *Order) error {
	sql, args, err := sq.Insert(order.TableName()).
		Columns("order_no", "user_id", "total_price", "postage", "status").
		Values(order.OrderNo, order.UserId, order.TotalPrice, order.Postage, order.Status).
		ToSql()

	if err != nil {
		return err
	}

	result, err := dao.DB().Exec(sql, args...)

	if err == nil {
		id, err := result.LastInsertId()
		if err == nil {
			order.ID = int(id)
		}
	}

	return err
}
