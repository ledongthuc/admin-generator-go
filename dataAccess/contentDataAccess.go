package dataAccess

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jbrodriguez/mlog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ledongthuc/admin-generator-go/helpers"
)

// TableDataAccess contains methods that used for access to `information_schema.tables`.
type contentDataAccess struct {
}

// Content is instance for ContentDataAccess
var Content contentDataAccess

// GetAll use to select all tables from table dynamically
func (dataAccess *contentDataAccess) GetAll(tableName string) []map[string]interface{} {
	configuration := helpers.LoadConfiguration()

	dbx, err := sqlx.Open(configuration.Type, configuration.ConnectionString)
	if err != nil {
		mlog.Error(err)
		return nil
	}

	query := `SELECT * FROM ` + strconv.Quote(tableName)
	rows, err := dbx.Queryx(query)
	dbx.Close()
	if err != nil {
		mlog.Error(err)
		return nil
	}

	var result []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		err = rows.MapScan(row)
		if err != nil {
			mlog.Error(err)
			continue
		}

		result = append(result, dataAccess.format(row))
	}

	return result
}

// GetAll use to select all tables from table dynamically
func (dataAccess *contentDataAccess) GetById(tableName string, idName string, id string) map[string]interface{} {
	configuration := helpers.LoadConfiguration()

	dbx, err := sqlx.Open(configuration.Type, configuration.ConnectionString)
	if err != nil {
		mlog.Error(err)
		return nil
	}

	query := `SELECT * FROM ` + strconv.Quote(tableName) + ` where ` + idName + ` = ` + id
	rows, err := dbx.Queryx(query)
	dbx.Close()
	if err != nil {
		mlog.Error(err)
		return nil
	}

	for rows.Next() {
		row := make(map[string]interface{})
		err = rows.MapScan(row)
		if err != nil {
			mlog.Error(err)
			continue
		}

		return dataAccess.format(row)
	}

	return nil
}

func (dataAccess *contentDataAccess) New(tableName string, data map[string]string) (int64, error) {
	configuration := helpers.LoadConfiguration()

	dbx, err := sqlx.Open(configuration.Type, configuration.ConnectionString)
	if err != nil {
		mlog.Error(err)
		return -1, err
	}

	var columns []string
	var values []string
	for column, value := range data {
		columns = append(columns, column)
		values = append(values, value)
	}
	whereClauseColumns := `"` + strings.Join(columns, "\",\"") + `"`
	whereClauseValues := `'` + strings.Join(values, "','") + `'`

	whereClause := fmt.Sprintf(`
            INSERT INTO
                %s (%s)
            VALUES
                (%s)
        `, tableName, whereClauseColumns, whereClauseValues)

	_, err = dbx.Exec(whereClause)
	if err != nil {
		return -1, err
	}

	return -1, nil
}

func (dataAccess *contentDataAccess) format(data map[string]interface{}) map[string]interface{} {
	for columnName, value := range data {
		switch value.(type) {
		case []uint8:
			data[columnName] = string(value.([]uint8))
			break
		case time.Time:
			data[columnName] = value.(time.Time).Format("2/1/2006 15:04:05")
			break
		}
	}
	return data
}
