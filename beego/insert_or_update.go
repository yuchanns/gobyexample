package beego

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID     int `orm:"pk"`
	Status int
}

func InsertOrUpdatePrintSql() error {
	db, mock, err := sqlmock.New()
	if err != nil {
		return err
	}
	defer db.Close()
	mock.ExpectPrepare("SELECT TIMEDIFF")
	mock.ExpectPrepare("SELECT ENGINE")
	mock.ExpectExec("INSERT").
		WillReturnResult(sqlmock.NewResult(1, 1))
	orm.Debug = true
	o, err := orm.NewOrmWithDB("mysql", "default", db)
	if err != nil {
		return err
	}
	_ = o.Using("db1")
	orm.RegisterModel(new(User))
	u := &User{
		ID:     1,
		Status: 0,
	}
	_, err = o.InsertOrUpdate(u)

	return err
}
