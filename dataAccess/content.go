package dataAccess

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jbrodriguez/mlog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ledongthuc/admin-generator-go/services"
)

// TableDataAccess contains methods that used for access to `information_schema.tables`.
type content struct {
}

// Content is instance for ContentDataAccess
var Content content

// GetAll use to select all tables from table dynamically
func (content *content) GetAll(tableName string) []map[string]interface{} {
	setting, err := services.Settings.Load()
	if err != nil {
		mlog.Error(err)
		return nil
	}

	dbx, err := sqlx.Open(setting.Database.Type, setting.Database.ConnectionString)
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

		result = append(result, content.format(row))
	}

	return result
}

// GetById get by ID
func (content *content) GetById(tableName string, idName string, id string) map[string]interface{} {
	setting, err := services.Settings.Load()
	if err != nil {
		mlog.Error(err)
		return nil
	}

	dbx, err := sqlx.Open(setting.Database.Type, setting.Database.ConnectionString)
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

		return content.format(row)
	}

	return nil
}

func (content *content) New(tableName string, data map[string]string) (int64, error) {
	setting, err := services.Settings.Load()
	if err != nil {
		mlog.Error(err)
		return -1, err
	}

	dbx, err := sqlx.Open(setting.Database.Type, setting.Database.ConnectionString)
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
	dbx.Close()
	if err != nil {
		return -1, err
	}

	return -1, nil
}

func (content *content) Update(tableName string, keyName string, keyValue string, data map[string]string) (int64, error) {
	if data == nil || len(data) <= 0 {
		return -1, errors.New("Data should be not empty")
	}

	setting, err := services.Settings.Load()
	if err != nil {
		mlog.Error(err)
		return -1, err
	}

	dbx, err := sqlx.Open(setting.Database.Type, setting.Database.ConnectionString)
	if err != nil {
		mlog.Error(err)
		return -1, err
	}

	var updatedParts []string
	for column, value := range data {
		updatedParts = append(updatedParts,
			fmt.Sprintf("%s = '%s'", column, value))
	}
	updatedClause := strings.Join(updatedParts, ", ")

	whereClause := fmt.Sprintf(`
            UPDATE
                %s
            SET
                %s
            WHERE
                %s = '%s'
        `, tableName, updatedClause, keyName, keyValue)
	mlog.Info(whereClause)
	_, err = dbx.Exec(whereClause)
	dbx.Close()
	if err != nil {
		return -1, err
	}

	return -1, nil
}

func (content *content) Delete(tableName string, keyName string, keyValue string) error {
	setting, err := services.Settings.Load()
	if err != nil {
		mlog.Error(err)
		return err
	}

	dbx, err := sqlx.Open(setting.Database.Type, setting.Database.ConnectionString)
	if err != nil {
		mlog.Error(err)
		return err
	}

	query := `DELETE FROM ` + strconv.Quote(tableName) + ` WHERE ` + keyName + ` = '` + keyValue + `'`
	mlog.Info("Run scripting: %s", query)
	_, err = dbx.Exec(query)
	return err
}

func (content *content) format(data map[string]interface{}) map[string]interface{} {
	for columnName, value := range data {
		switch value.(type) {
		case []uint8:
			data[columnName] = string(value.([]uint8))
			break
		}
	}
	return data
}
