package squirrel

type Order struct {
	ID         int    `db:"id" json:"id"`
	OrderNo    string `db:"order_no" json:"order_no"`
	UserId     int    `db:"user_id" json:"user_id"`
	TotalPrice int    `db:"total_price" json:"total_price"`
	Postage    int    `db:"postage" json:"postage"`
	Status     int    `db:"status" json:"status"`
	IsDeleted  int    `db:"is_deleted" json:"is_deleted"`
	CreatedAt  int64  `db:"created_at" json:"-"`
	UpdatedAt  int64  `db:"updated_at" json:"-"`
	DeletedAt  int64  `db:"deleted_at" json:"-"`
}

func (Order) TableName() string {
	return "`order`"
}
