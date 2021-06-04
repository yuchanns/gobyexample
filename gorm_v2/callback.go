package gorm_v2

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func username(db *gorm.DB) string {
	userName := ""
	if userNameIFace, ok := db.InstanceGet("username"); ok && userNameIFace != nil {
		if s, ok := userNameIFace.(string); ok {
			userName = s
		}
	}
	return userName
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

type columns []*struct {
	Name  string
	Value interface{}
}

// deleteCallback
// checkout about callback in https://gorm.io/docs/write_plugins.html
// change to update if soft delete is specified
// add sql clauses by db.Statement.AddClause
// then build a sql string by db.Statement.Build
func deleteCallback(db *gorm.DB) {
	if db.Error != nil || db.Statement.Schema == nil {
		return
	}
	var extraOption string
	if str, ok := db.Statement.Get("gorm:delete_option"); ok {
		extraOption = fmt.Sprint(str)
	}
	deletedStateField := db.Statement.Schema.LookUpField("DeletedState")
	// hard delete
	if db.Statement.Unscoped || deletedStateField == nil {
		db.Statement.AddClause(clause.Delete{})
		db.Statement.AddClause(clause.From{Tables: []clause.Table{{Name: db.Statement.Table}}})
		db.Statement.Build(db.Callback().Delete().Clauses...)
		db.Exec(fmt.Sprintf("%v%v", db.Statement.SQL.String(), addExtraSpaceIfExist(extraOption)))
		return
	}
	// soft delete
	userName := username(db)
	sets := clause.Set{clause.Assignment{Column: clause.Column{Name: deletedStateField.DBName}, Value: 1}}
	now := time.Now()
	cols := columns{
		{Name: "DeletedOn", Value: now}, {Name: "DeletedBy", Value: userName},
		{Name: "ModifiedOn", Value: now}, {Name: "ModifiedBy", Value: userName},
	}
	for _, col := range cols {
		if field := db.Statement.Schema.LookUpField(col.Name); field != nil {
			sets = append(sets, clause.Assignment{Column: clause.Column{Name: field.DBName}, Value: col.Value})
		}
	}
	db.Statement.AddClause(sets)
	db.Statement.AddClause(clause.Update{Table: clause.Table{Name: db.Statement.Table}})
	db.Statement.Build(db.Callback().Update().Clauses...)
	db.Exec(fmt.Sprintf("%v%v", db.Statement.SQL.String(), addExtraSpaceIfExist(extraOption)))
}

// updateTimeStampForCreateCallback
// checkout about callback in https://gorm.io/docs/write_plugins.html
// auto set values if fields about create exist
func updateTimeStampForCreateCallback(db *gorm.DB) {
	if db.Error != nil || db.Statement.Schema == nil {
		return
	}
	userName := username(db)
	now := time.Now()
	cols := columns{
		{Name: "CreatedOn", Value: now}, {Name: "CreatedBy", Value: userName},
		{Name: "ModifiedOn", Value: now}, {Name: "ModifiedBy", Value: userName},
	}
	for _, col := range cols {
		if field := db.Statement.Schema.LookUpField(col.Name); field != nil {
			if _, isZero := field.ValueOf(db.Statement.ReflectValue); isZero {
				_ = field.Set(db.Statement.ReflectValue, col.Value)
			}
		}
	}
}

// updateTimeStampForUpdateCallback
// checkout about callback in https://gorm.io/docs/write_plugins.html
// auto set values if fields about update exist
func updateTimeStampForUpdateCallback(db *gorm.DB) {
	if db.Error != nil || db.Statement.Schema == nil {
		return
	}
	// won't update without model changed
	if !db.Statement.Changed() {
		return
	}
	db.Statement.Omit("created_on", "created_by", "deleted_state", "deleted_by", "createdOn", "createdBy", "deletedState", "deletedBy")
	userName := username(db)
	now := time.Now()
	cols := columns{
		{Name: "ModifiedOn", Value: now}, {Name: "ModifiedBy", Value: userName},
	}
	for _, col := range cols {
		if field := db.Statement.Schema.LookUpField(col.Name); field != nil {
			db.Statement.SetColumn(col.Name, col.Value)
		}
	}
}

// softDeleteQueryCallback
// checkout about callback in https://gorm.io/docs/write_plugins.html
// add deleted_state = 0 while performing a query
func softDeleteQueryCallback(db *gorm.DB) {
	if db.Error != nil || db.Statement.Unscoped || db.Statement.Schema == nil {
		return
	}
	if deletedField := db.Statement.Schema.LookUpField("DeletedState"); deletedField != nil {
		column := clause.Column{Name: deletedField.DBName}
		// add table name as column prefix if join
		if len(db.Statement.Joins) > 0 {
			column.Table = db.Statement.Schema.Table
		}
		db.Statement.AddClause(clause.Where{Exprs: []clause.Expression{clause.Eq{Column: column, Value: 0}}})
	}
}
