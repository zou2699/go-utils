/*
@Time : 2021/3/20 13:43
@Author : Tux
@Description :
*/

package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库的连接核心
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

// 数据库的连接基本信息
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

// 存储column信息
type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

// 生成 DBModel
func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

// 连接数据库
func (m *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local",
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	err = m.DBEngine.Ping()
	if err != nil {
		return err
	}
	return err
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT " +
		"COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(
			&column.ColumnName,
			&column.DataType,
			&column.ColumnKey,
			&column.IsNullable,
			&column.ColumnType,
			&column.ColumnComment,
		)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}
	return columns, nil
}

// DataType字段的类型与go类型的映射表
var dBTypeToStructTypeMap = map[string]string{
	"int":          "int32",
	"tinyint":      "int8",
	"smallint":     "int",
	"mediumint":    "int64",
	"bigint":       "int64",
	"bit":          "int",
	"bool":         "bool",
	"enum":         "string",
	"set":          "string",
	// "varchar":      "string",
	// "varchar(128)": "string",
	// "varchar(255)": "string",
	"timestamp":    "time.time",
}

func DBTypeToStructType(key string) string {
	if strings.Contains(key, "varchar") {
		return "string"
	}
	if v, ok := dBTypeToStructTypeMap[key]; ok {
		return v
	}
	return ""
}
