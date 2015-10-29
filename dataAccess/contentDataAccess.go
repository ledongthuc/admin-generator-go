package dataAccess

import (
	"log"
	"strconv"
	"time"

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
		log.Println(err)
		return nil
	}

	query := `SELECT * FROM ` + strconv.Quote(tableName)
	rows, err := dbx.Queryx(query)
	dbx.Close()
	if err != nil {
		log.Println(err)
		return nil
	}

	var result []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		err = rows.MapScan(row)
		if err != nil {
			log.Println(err)
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
		log.Println(err)
		return nil
	}

	query := `SELECT * FROM ` + strconv.Quote(tableName) + ` where ` + idName + ` = ` + id
	rows, err := dbx.Queryx(query)
	dbx.Close()
	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {
		row := make(map[string]interface{})
		err = rows.MapScan(row)
		if err != nil {
			log.Println(err)
			continue
		}

		return dataAccess.format(row)
	}

	return nil
}

func (dataAccess *contentDataAccess) format(data map[string]interface{}) map[string]interface{} {
	for columnName, value := range data {
		switch valueType := value.(type) {
		case []uint8:
			data[columnName] = string(value.([]uint8))
		case time.Time:
			data[columnName] = value.(time.Time).Format("2/1/2006 15:04:05")
		default:
			log.Println(valueType)
		}
	}
	return data
}
