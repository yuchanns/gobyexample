package beego

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID     int `orm:"pk"`
	Status int
}

type TExchangeInfo struct {
	ID           int64     `orm:"column(id);pk"`
	DeparmentID  int64     `orm:"column(deparment_id)"`
	Times        uint      `orm:"column(times)"`
	Number       uint      `orm:"column(number)"`
	Lastmodified time.Time `orm:"column(lastmodified);type(datetime);auto_now"`
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
	orm.RegisterModel(new(TExchangeInfo))
	u := &TExchangeInfo{
		ID:          10086,
		DeparmentID: 1,
		Times:       0,
		Number:      10,
	}
	_, err = o.InsertOrUpdate(u)

	return err
}
