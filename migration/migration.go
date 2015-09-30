package migration

import (
    "fmt"
    "strings"

    "github.com/ledongthuc/admin-generator-go/database"
)

func Run() {
    tableDataAccess := database.NewTableDataAccess()
    tables := tableDataAccess.GetAll()

    var result []string
    for _, table := range tables {
        result = append(result, table.Name)

    }
    fmt.Println(strings.Join(result, ","))

    return
}
