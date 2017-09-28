package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
	"os"
)


func InitDatabase() *sql.DB {
	//删除原数据 不对数据做持久化
	db_path := "/tmp/sqlite.db"
	os.Remove(db_path)
	//新建数据文件
	db,err := sql.Open("sqlite3",db_path)
	Check(err)
	return db
}

func CreateTable(db *sql.DB,ranged_key []string)  {
	//根据解析的日志格式创建log表
	//构造创建sql
	sqlstmt := "create table log ("
	for i,key := range ranged_key {
		if i == len(ranged_key) - 1 {
			switch key {
			case "request_time":
				sqlstmt = sqlstmt + key + " float"
			case "body_bytes_sent":
				sqlstmt = sqlstmt + key + " integer"
			default:
				sqlstmt = sqlstmt + key + " text"
			}
		} else {
			switch key {
			case "request_time":
				sqlstmt = sqlstmt + key + " float,"
			case "body_bytes_sent":
				sqlstmt = sqlstmt + key + " integer,"
			default:
				sqlstmt = sqlstmt + key + " text,"
			}
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
	insert_column := ""
	insert_value := ""
	// 构造插入语句
	for i,key := range column {
		if i != len(column) - 1 {
			insert_column = insert_column + key + ","
			insert_value = insert_value + "?,"
		} else {
			insert_column = insert_column + key
			insert_value = insert_value + "?"
		}
	}
	stmt, err := db.Prepare("insert into log(" + insert_column + ") values(" + insert_value + ")")
	Check(err)
	//不定参数 类型转化
	arg := make([]interface{},len(value))
	for i,v := range value {
		arg[i] = v
	}
	_, err = stmt.Exec(arg...)
	Check(err)
}

func CountUniqueIP(db *sql.DB) string {
	res, err := db.Query("SELECT COUNT(DISTINCT remote_addr) AS count FROM log")
	Check(err)
	defer res.Close()
	var count string
	for res.Next() {
		res.Scan(&count)
	}
	//fmt.Println(count)
	return count
}

func CountRequest(db *sql.DB) string {
	res, err := db.Query("SELECT COUNT(*) AS count FROM log")
	Check(err)
	defer res.Close()
	var count string
	for res.Next() {
		res.Scan(&count)
	}
	//fmt.Println(count)
	return count
}

func ListPopularURL(db *sql.DB) [200][2]string {
	res, err := db.Query("SELECT request,count(*) AS count FROM log GROUP BY request ORDER BY count DESC LIMIT 200")
	Check(err)
	defer res.Close()
	var rows [200][2]string
	index := 0
	for res.Next() {
		res.Scan(&rows[index][0],&rows[index][1])
		index++
	}
	return rows

}

func ListPopularIP(db *sql.DB) [50][2]string {
	res, err := db.Query("SELECT remote_addr,count(*) AS count FROM log GROUP BY remote_addr ORDER BY count DESC LIMIT 50")
	Check(err)
	defer res.Close()
	var rows [50][2]string
	index := 0
	for res.Next() {
		res.Scan(&rows[index][0],&rows[index][1])
		index++
	}
	return rows

}




