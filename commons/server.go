package commons

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

//type server struct {
//	id int
//	servername string
//	user string
//	password string
//	created_at string
//	updated_at string
//	deleted_at string
//}

func findServer() {
	//conn, _ := client.Connect(log.Debugf, 100, 400, 5, "127.0.0.1:3306", `db_user`, `db_password`, `echo_db`)
	db, err := sql.Open("mysql", "db_user:db_password@localhost/automation_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM server")
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
}


