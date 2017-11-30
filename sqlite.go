package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
	"strconv"
	"os"
	"strings"
)

func InitDatabase() *sql.DB {
	db_path := "/tmp/sqlite.db"
	//判断是否要重新读入数据
	if t.TruncateDatabase {
		os.Remove(db_path)
	}
	// 新建数据文件
	db, err := sql.Open("sqlite3", db_path)
	Check(err)
	return db
}

// 建表
func CreateTable(db *sql.DB, ranged_key []string) {
	// 根据解析的日志格式创建log表
	// 构造创建sql
	sqlstmt := "create table log ("
	for i, key := range ranged_key {
		if i == len(ranged_key)-1 {
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

func InsertData(db *sql.DB, column []string, value []string) {
	insert_column := ""
	insert_value := ""
	// 构造插入语句
	for i, key := range column {
		if i != len(column)-1 {
			insert_column = insert_column + key + ","
			insert_value = insert_value + "?,"
		} else {
			insert_column = insert_column + key
			insert_value = insert_value + "?"
		}
	}
	stmt, err := db.Prepare("insert into log(" + insert_column + ") values(" + insert_value + ")")
	Check(err)
	// 不定参数 类型转化
	arg := make([]interface{}, len(value))
	for i, v := range value {
		arg[i] = v
	}
	_, err = stmt.Exec(arg...)
	Check(err)
}

// 渲染长度为2的数组
func RenderTwoColumn(db *sql.DB, sql string) [][2]string {
	res, err := db.Query(sql)
	Check(err)
	defer res.Close()
	var rows [][2]string
	for res.Next() {
		var a [2]string
		res.Scan(&a[0], &a[1])
		rows = append(rows, a)
	}
	return rows
}

// 渲染结果字符串
func RenderString(db *sql.DB, sql string) string {
	res, err := db.Query(sql)
	Check(err)
	defer res.Close()
	var result string
	for res.Next() {
		res.Scan(&result)
	}
	return result
}

// 统计独立ip数
func CountUniqueIP(db *sql.DB) string {
	sql := "SELECT COUNT(DISTINCT remote_addr) AS count FROM log"
	count := RenderString(db, sql)
	return count
}

// 统计请求数
func CountRequest(db *sql.DB) string {
	sql := "SELECT COUNT(*) AS count FROM log"
	count := RenderString(db, sql)
	return count
}

// 统计访问量前xx的请求
func ListPopularURL(db *sql.DB) [][2]string {
	sql := "SELECT count(*) AS count,request FROM log GROUP BY request ORDER BY count DESC LIMIT " + t.TopRequest
	rows := RenderTwoColumn(db, sql)
	return rows
}

// 统计访问量前xx的IP
func ListPopularIP(db *sql.DB) [][2]string {
	sql := "SELECT count(*) AS count,remote_addr FROM log GROUP BY remote_addr ORDER BY count DESC LIMIT " + t.TopIP
	rows := RenderTwoColumn(db, sql)
	return rows
}

// 统计平均超时请求
func ListAvg(db *sql.DB) [][2]string {
	sql := "SELECT * FROM (SELECT round(avg(request_time),3) AS cost,request FROM log  GROUP BY request ORDER BY cost DESC) WHERE cost > " + t.Overtime
	rows := RenderTwoColumn(db, sql)
	return rows
}

// 统计最长超时请求
func ListLongest(db *sql.DB) [][6]string {
	sql := "SELECT * FROM (SELECT request_time,request,remote_addr,time_local,http_referer FROM log GROUP BY request ORDER BY max(request_time) DESC) WHERE request_time > " + t.Overtime
	res, err := db.Query(sql)
	Check(err)
	defer res.Close()
	var rows [][6]string
	for res.Next() {
		var a [6]string
		res.Scan(&a[1], &a[2], &a[3], &a[4], &a[5])
		// 将请求方法截去
		tmp_request := strings.Split(a[2], " ")
		request := strings.Split(tmp_request[1], "/")
		key := len(request)
		// 匹配接口
		a[0] = t.DeveloperMap[request[key-1]]
		// 接口没有匹配到则匹配控制器
		if a[0] == "" && key > 3 {
			a[0] = t.DeveloperMap[request[2]]
		}
		rows = append(rows, a)
	}
	return rows
}

// 统计异常请求
func ListError(db *sql.DB) ([][2]string, string) {
	sql := "SELECT * FROM (SELECT status,request FROM log GROUP BY status,request) WHERE status LIKE '4%' OR status LIKE '5%'"
	rows := RenderTwoColumn(db, sql)
	sql = "SELECT count(*) AS count FROM log WHERE status LIKE '4%' OR status LIKE '5%'"
	count, _ := strconv.ParseFloat(RenderString(db, sql), 6)
	total, _ := strconv.ParseFloat(CountRequest(db), 6)
	rate := count / total * 100
	s := fmt.Sprintf("%.3f", rate)
	return rows, s
}

// 统计时间段访问量
func CountByTime(db *sql.DB) [][2]string {
	sql := "SELECT count(*) AS count,time_local FROM log GROUP BY time_local"
	rows := RenderTwoColumn(db, sql)
	return rows
}
