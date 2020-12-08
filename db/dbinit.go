package db

import (
	"database/sql"
	"log"
)

func DBinit(database *sql.DB) {
	DDLs := []string{CreatingUsersTable, CreatingATMsTable, CreatingAccountsTable, CreatingTransactionsHistoryTable}
	for _, ddl := range DDLs {
		_, err := database.Exec(ddl)
		if err != nil {
			log.Fatalf("Can't init %s error is %d", ddl, err)
		}
	}
}
