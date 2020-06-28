package db

import (
	"database/sql"
	"errors"
	log "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func (p *DbPool) Open() error {
	var err error
	p.Db.SQLDB, err = sql.Open(p.Db.DriverName, p.Db.DataSourceName)
	if err != nil {
		return err
	}
	if err = p.Db.SQLDB.Ping(); err != nil {
		return err
	}
	p.Db.SQLDB.SetMaxOpenConns(p.Db.MaxOpenConns)
	p.Db.SQLDB.SetMaxIdleConns(p.Db.MaxIdleConns)
	return nil
}

func (p *DbPool) Close() error {
	return p.Db.SQLDB.Close()
}

func (p *DbPool) Get(queryStr string, args ...interface{}) (map[string]interface{}, error) {
	results, err := p.Query(queryStr, args...)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if len(results) <= 0 {
		return map[string]interface{}{}, sql.ErrNoRows
	}
	if len(results) > 1 {
		return map[string]interface{}{}, errors.New("sql: more than one rows")
	}
	return results[0], nil
}

func (p *DbPool) Query(queryStr string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := p.Db.SQLDB.Query(queryStr, args...)
	if err != nil {
		log.Error(err)
		return []map[string]interface{}{}, err
	}
	defer rows.Close()
	columns, err := rows.ColumnTypes()
	scanArgs := make([]interface{}, len(columns))
	values := make([]sql.RawBytes, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	rowsMap := make([]map[string]interface{}, 0, 10)
	for rows.Next() {
		rows.Scan(scanArgs...)
		rowMap := make(map[string]interface{})
		for i, value := range values {
			rowMap[columns[i].Name()] = bytes2RealType(value, columns[i])
		}
		rowsMap = append(rowsMap, rowMap)
	}
	if err = rows.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return rowsMap, nil
}

func (p *DbPool) Exec(sqlStr string, args ...interface{}) (sql.Result, error) {
	return p.Db.SQLDB.Exec(sqlStr, args...)
}

func (p *DbPool) Update(updateStr string, args ...interface{}) (int64, error) {
	result, err := p.Exec(updateStr, args...)
	if err != nil {
		return 0, err
	}
	affect, err := result.RowsAffected()
	return affect, err
}

func (p *DbPool) Insert(insertStr string, args ...interface{}) (int64, error) {
	result, err := p.Exec(insertStr, args...)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	return lastId, err

}

func (p *DbPool) Delete(deleteStr string, args ...interface{}) (int64, error) {
	result, err := p.Exec(deleteStr, args...)
	if err != nil {
		return 0, err
	}
	affect, err := result.RowsAffected()
	return affect, err
}

type SQLConnTransaction struct {
	SQLTX *sql.Tx
}

func (p *DbPool) Begin() (*SQLConnTransaction, error) {
	var oneSQLConnTransaction = &SQLConnTransaction{}
	var err error
	if pingErr := p.Db.SQLDB.Ping(); pingErr == nil {
		oneSQLConnTransaction.SQLTX, err = p.Db.SQLDB.Begin()
	}
	return oneSQLConnTransaction, err
}

func (t *SQLConnTransaction) Rollback() error {
	return t.SQLTX.Rollback()
}

func (t *SQLConnTransaction) Commit() error {
	return t.SQLTX.Commit()
}

func (t *SQLConnTransaction) Get(queryStr string, args ...interface{}) (map[string]interface{}, error) {
	results, err := t.Query(queryStr, args...)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if len(results) <= 0 {
		return map[string]interface{}{}, sql.ErrNoRows
	}
	if len(results) > 1 {
		return map[string]interface{}{}, errors.New("sql: more than one rows")
	}
	return results[0], nil
}

func (t *SQLConnTransaction) Query(queryStr string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := t.SQLTX.Query(queryStr, args...)
	if err != nil {
		log.Error(err)
		return []map[string]interface{}{}, err
	}
	defer rows.Close()
	columns, err := rows.ColumnTypes()
	scanArgs := make([]interface{}, len(columns))
	values := make([]sql.RawBytes, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	rowsMap := make([]map[string]interface{}, 0, 10)
	for rows.Next() {
		rows.Scan(scanArgs...)
		rowMap := make(map[string]interface{})
		for i, value := range values {
			rowMap[columns[i].Name()] = bytes2RealType(value, columns[i])
		}
		rowsMap = append(rowsMap, rowMap)
	}
	if err = rows.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return rowsMap, nil
}

func (t *SQLConnTransaction) Exec(sqlStr string, args ...interface{}) (sql.Result, error) {
	return t.SQLTX.Exec(sqlStr, args...)
}

func (t *SQLConnTransaction) Update(updateStr string, args ...interface{}) (int64, error) {
	result, err := t.Exec(updateStr, args...)
	if err != nil {
		return 0, err
	}
	affect, err := result.RowsAffected()
	return affect, err
}

func (t *SQLConnTransaction) Insert(insertStr string, args ...interface{}) (int64, error) {
	result, err := t.Exec(insertStr, args...)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	return lastId, err

}

func (t *SQLConnTransaction) Delete(deleteStr string, args ...interface{}) (int64, error) {
	result, err := t.Exec(deleteStr, args...)
	if err != nil {
		return 0, err
	}
	affect, err := result.RowsAffected()
	return affect, err
}

func bytes2RealType(src []byte, column *sql.ColumnType) interface{} {
	srcStr := string(src)
	var result interface{}
	switch column.DatabaseTypeName() {
	case "BIT", "TINYINT", "SMALLINT", "INT":
		result, _ = strconv.ParseInt(srcStr, 10, 64)
	case "BIGINT":
		result, _ = strconv.ParseUint(srcStr, 10, 64)
	case "CHAR", "VARCHAR",
		"TINY TEXT", "TEXT", "MEDIUM TEXT", "LONG TEXT",
		"TINY BLOB", "MEDIUM BLOB", "BLOB", "LONG BLOB",
		"JSON", "ENUM", "SET",
		"YEAR", "DATE", "TIME", "TIMESTAMP", "DATETIME":
		result = srcStr
	case "FLOAT", "DOUBLE", "DECIMAL":
		result, _ = strconv.ParseFloat(srcStr, 64)
	default:
		result = nil
	}
	return result
}