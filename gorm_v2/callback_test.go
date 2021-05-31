package gorm_v2

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type User struct {
	ID           int
	Name         string
	Gender       int
	DeletedState int
	CreatedBy    string
	CreatedOn    time.Time
	ModifiedBy   string
	ModifiedOn   time.Time
	DeletedBy    string
	DeletedOn    time.Time
}

type Book struct {
	ID   int
	Name string
}

func TestCallback_Delete(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.Nil(t, err)

	assert.Nil(t, db.Callback().Delete().Replace("gorm:delete", deleteCallback))

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users` SET `deleted_state`=(.+),`deleted_on`=(.+),`deleted_by`=(.+),`modified_on`=(.+),`modified_by`=(.+) WHERE id = (.+) and gender = (.+) OPTION (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `books` WHERE id = (.+) and gender = (.+) OPTION (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	db.InstanceSet("username", "yuchanns").Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&User{}, "id = ? and gender = ?", 1, 2)
	db.InstanceSet("username", "yuchanns").Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&Book{}, "id = ? and gender = ?", 1, 2)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCallback_UpdateTimestampForCreate(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.Nil(t, err)

	assert.Nil(t, db.Callback().Create().Replace("gorm:before_create", updateTimeStampForCreateCallback))

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` \\(`name`,`gender`,`deleted_state`,`created_by`,`created_on`,`modified_by`,`modified_on`,`deleted_by`,`deleted_on`\\) VALUES (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` \\(`name`,`gender`,`deleted_state`,`created_by`,`created_on`,`modified_by`,`modified_on`,`deleted_by`,`deleted_on`\\) VALUES (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `books`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	db.InstanceSet("username", "yuchanns").Create(&User{
		Name:   "yuchanns",
		Gender: 1,
	})
	db.InstanceSet("username", "yuchanns").Create(&User{
		Name:      "yuchanns",
		Gender:    1,
		CreatedOn: time.Date(2020, time.May, 27, 11, 26, 0, 0, time.Local),
	})
	db.InstanceSet("username", "yuchanns").Create(&Book{
		Name: "yuchanns",
	})
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCallback_UpdateTimeStampBeforeUpdate(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.Nil(t, err)

	assert.Nil(t, db.Callback().Update().Replace("gorm:before_update", updateTimeStampForUpdateCallback))

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users` SET `modified_by`=(.+),`modified_on`=(.+),`name`=(.+) WHERE `id` = (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `books` SET `name`=(.+) WHERE `id` = (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	fixedTime := time.Date(2020, time.May, 27, 11, 26, 0, 0, time.Local)
	db.InstanceSet("username", "yuchanns").Model(&User{
		ID: 1, Name: "yuchanns", Gender: 1,
		CreatedOn: fixedTime, CreatedBy: "yuchanns1",
		ModifiedOn: fixedTime, ModifiedBy: "yuchanns1",
	}).Update("name", "yuchanns1")
	db.InstanceSet("username", "yuchanns").Model(&User{
		ID: 1, Name: "yuchanns", Gender: 1,
		CreatedOn: fixedTime, CreatedBy: "yuchanns1",
		ModifiedOn: fixedTime, ModifiedBy: "yuchanns1",
	}).Omit("not_exist").Update("not_exist", "")
	db.InstanceSet("username", "yuchanns").Model(&Book{
		ID: 1, Name: "yuchanns1",
	}).Update("name", "yuchanns1")
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCallback_SoftDeleteQuery(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.Nil(t, err)

	assert.Nil(t, db.Callback().Row().Before("gorm:row").Register("soft_delete_query", softDeleteQueryCallback))
	assert.Nil(t, db.Callback().Query().Before("gorm:query").Register("soft_delete_query", softDeleteQueryCallback))

	now := time.Now()
	rows := sqlmock.NewRows([]string{
		"id", "name", "gender", "deleted_state", "created_on", "created_by",
		"modified_on", "modified_by", "deleted_by", "deleted_on",
	})
	rows.AddRow(1, "yuchanns", 1, 0, now, "yuchanns", now, "yuchanns", "", nil)
	mock.ExpectQuery("SELECT (.+) FROM `users` WHERE name like (.+)  AND `deleted_state` = (.+)").WillReturnRows(rows)
	mock.ExpectQuery("SELECT (.+) FROM `users` WHERE name like (.+)  AND `deleted_state` = (.+)").WillReturnRows(rows)
	_, _ = db.InstanceSet("username", "yuchanns").Model(&User{}).Where("name like ? ", "%yuchanns%").Rows()
	var users []*User
	db.InstanceSet("username", "yuchanns").Where("name like ? ", "%yuchanns%").Find(&users)
	assert.Nil(t, mock.ExpectationsWereMet())
}
