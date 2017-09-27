package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"log"
	"fmt"
)

func InitDatabase() *sql.DB {
	//删除原数据 不对数据做持久化
	os.Remove("./sqlite.db")
	//新建数据文件
	db,err := sql.Open("sqlite3","./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateTable(db *sql.DB,ranged_key []string)  {
	//根据解析的日志格式创建log表
	//构造创建sql
	sqlstmt := "create table log ("
	for i,key := range ranged_key {
		if i == len(ranged_key) - 1 {
			sqlstmt = sqlstmt + key + " text"
		} else {
			sqlstmt = sqlstmt + key + " text, "
		}
	}
	sqlstmt = sqlstmt + ");"
	fmt.Println(sqlstmt)
	_, err := db.Exec(sqlstmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlstmt)
	}
}

func InsertData(db *sql.DB, column []string, value []string)  {
	// 事务的写法
	//tx, err := db.Begin()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer stmt.Close()
	//for i := 0; i < 100; i++ {
	//	_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//tx.Commit()
	insert_column := ""
	insert_value := ""
	for i,key := range column {
		if i != len(column) - 1 {
			insert_column = insert_column + key + ","
			insert_value = insert_value + "?,"
		} else {
			insert_column = insert_column + key
			insert_value = insert_value + "?"
		}
	}
	fmt.Println(insert_column,insert_value)
	stmt, err := db.Prepare("insert into log(" + insert_column + ") values(" + insert_value + ")")
	if err != nil {
		log.Fatal(err)
	}
	//不定参数 类型转化
	arg := make([]interface{},len(value))
	for i,v := range value {
		arg[i] = v
	}
	_, err = stmt.Exec(arg...)
	if err != nil {
		log.Fatal(err)
	}
}

//func QueryData(db *sql.DB)  {
//	rows, err := db.Query("select id, name from foo")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var id int
//		var name string
//		err = rows.Scan(&id, &name)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(id, name)
//	}
//	err = rows.Err()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	stmt, err = db.Prepare("select name from foo where id = ?")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer stmt.Close()
//	var name string
//	err = stmt.QueryRow("3").Scan(&name)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(name)
//}
//





