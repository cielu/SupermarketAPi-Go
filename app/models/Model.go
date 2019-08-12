package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"reflect"
	"strings"
	"supermarket/library/setting"
)

var db *gorm.DB

type Model struct {
	CreatedAt FormatTime `json:"created_at"`
	UpdatedAt FormatTime `json:"updated_at"`
}

// func InsertFile() {
// 	jsonData, err := ioutil.ReadFile("runtime/app/gussULikeTab.json")
// 	if err != nil {
// 		log.Fatal("读取文件失败", err);
// 		return
// 	}
// 	err = json.Unmarshal(jsonData, &gussULikeTabArr)
// 	if err != nil {
// 		log.Fatal("读取数据失败: ", err);
// 		return
// 	}
// 	tabData := make([]interface{}, len(gussULikeTabArr))
// 	for i, v := range gussULikeTabArr {
// 		tabData[i] = v
// 	}
// 	BatchInsert(tabData)
// 	return
// }

/*
 * 批量插入
 */
func BatchInsert(objArr []interface{}) (int64, error) {
	// If there is no data, nothing to do.
	if len(objArr) == 0 {
		return 0, errors.New("insert a slice length of 0")
	}

	mainObj := objArr[0]
	mainScope := db.NewScope(mainObj)
	mainFields := mainScope.Fields()
	quoted := make([]string, 0, len(mainFields))

	for i := range mainFields {
		// 主键未传入数据或传0的，或者需要忽略的，直接忽略
		if (mainFields[i].IsPrimaryKey && mainFields[i].IsBlank) || (mainFields[i].IsIgnored) {
			continue
		}
		quoted = append(quoted, mainScope.Quote(mainFields[i].DBName))
	}

	placeholdersArr := make([]string, 0, len(objArr))

	for _, obj := range objArr {
		scope := db.NewScope(obj)
		fields := scope.Fields()
		placeholders := make([]string, 0, len(fields))
		for i := range fields {
			if (fields[i].IsPrimaryKey && fields[i].IsBlank) || (fields[i].IsIgnored) {
				continue
			}
			vars := fields[i].Field.Interface()

			placeholders = append(placeholders, scope.AddToVars(vars))
		}
		placeholdersStr := "(" + strings.Join(placeholders, ", ") + ")"
		placeholdersArr = append(placeholdersArr, placeholdersStr)
		// add real variables for the replacement of placeholders' '?' letter later.
		mainScope.SQLVars = append(mainScope.SQLVars, scope.SQLVars...)
	}
	mainScope.Raw(fmt.Sprintf("INSERT INTO %s (%s) VALUES %s",
		mainScope.QuotedTableName(),
		strings.Join(quoted, ", "),
		strings.Join(placeholdersArr, ", "),
	))
	// Execute and Log
	if err := mainScope.Exec().DB().Error; err != nil {
		return 0, err
	}
	return mainScope.DB().RowsAffected, nil
}

// 分页
func Pagination(sql *gorm.DB, where map[string]interface{}, dataPtr interface{}) (mapData map[string]interface{}) {
	// page
	var (
		curPage  uint = 1
		pageSize      = setting.AppCfg.PageSize
	)
	mapData = make(map[string]interface{})
	// 分页大小
	if length, ok := where["pageSize"]; ok {
		pageSize = length.(uint)
	}
	// 获取传过来的分页，并且分页大于第一页
	if page, ok := where["page"]; ok && page != 1 {
		curPage = page.(uint)
	}
	// 获取数据
	sql.Offset(curPage * pageSize).Limit(pageSize).Find(dataPtr)
	// 通过反射获取数据
	dataValue := reflect.ValueOf(dataPtr).Elem()
	// 数据
	mapData["data"] = dataValue.Interface()
	mapData["curPage"] = curPage
	mapData["curPageSize"] = pageSize
	mapData["curTotal"] = dataValue.Len()
	// map data
	return
}

// 初始化数据库连接等
func Initialized() {
	// 申明变量
	var (
		err error
		connectSql string
		databaseCfg = setting.DatabaseCfg
	)
	switch databaseCfg.DbType {
	case "mysql":
		connectSql = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local",
			databaseCfg.User, databaseCfg.Password, databaseCfg.Host+":"+databaseCfg.Port, databaseCfg.DbName)
		break
	case "postgres":
		connectSql = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			databaseCfg.Host, databaseCfg.Port, databaseCfg.User, databaseCfg.DbName, databaseCfg.Password)
		break
	case "sqlite3":
		connectSql = "runtime/app/supermarket.db"
		break
	case "mssql":
		connectSql = fmt.Sprintf("sqlserver://%s:%s@tcp(%s)?database=%s",
			databaseCfg.User, databaseCfg.Password, databaseCfg.Host+":"+databaseCfg.Port, databaseCfg.DbName)
		break
	}
	// 连接数据库
	db, err = gorm.Open(databaseCfg.DbType, connectSql)
	if err != nil {
		log.Fatalf("database connection error: %v", err);
		return
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return databaseCfg.TablePrefix + defaultTableName
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Callback().Query()
	// debug 模式开启sql日志
	if setting.AppCfg.AppMode == "debug" {
		db.LogMode(true)
	}
}

// 关闭数据库连接
func CloseDb() {
	db.Close()
}
