package mysqlutil

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// confMySQL := config.C.MySQL
	// strConn := confMySQL.User + ":" + confMySQL.Password + "@tcp(" + confMySQL.Host + ":" + confMySQL.Port + ")/" + confMySQL.Dbname
	// strParameter := "?charset=" + confMySQL.Charset
	// strConn = strConn + strParameter
	// fmt.Println("Conn string: ", strConn)
	// db, _ = sql.Open("mysql", strConn)
	// db.SetMaxOpenConns(2000)
	// db.SetMaxIdleConns(1000)
	// db.Ping()
}

func Query(sql string) {
	rows, err := db.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Printf("Query MySQL DB Failed: %v", err)
	}
	for rows.Next() {
		var id int
		var title string
		var content string
		err = rows.Scan(&id, &title, &content)
		fmt.Println(id)
		fmt.Println(title)
		fmt.Println(content)
	}
}
