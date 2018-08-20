package driver

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func mysqlBuildDSN(config Config) string {
	c := mysql.NewConfig()
	c.User = "root"
	c.Passwd = ""
	c.Net = "tcp"
	c.Addr = "127.0.0.1:3306"
	c.DBName = config.DbName
	return c.FormatDSN()
}

func (d *Database) mysqlTableNames() ([]string, error) {
	rows, err := d.db.Query("show tables")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := []string{}
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func (d *Database) mysqlDumpTableDDL(table string) (string, error) {
	var ddl string
	sql := fmt.Sprintf("show create table %s;", table) // TODO: escape table name

	err := d.db.QueryRow(sql).Scan(&table, &ddl)
	if err != nil {
		return "", err
	}

	return ddl, nil
}
