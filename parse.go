package main

import (
	"os/exec"
	"bytes"
	"fmt"
	"strings"
	"database/sql"
	"os"
	"bufio"
)

// 配置文件结构体
type Conf struct {
	InputFile string      // 需要分析的日志文件路径
	StartDate string      // 开始日期
	EndDate string        // 结束日期
	Overtime string       // 超时时间
	TopIP string
	TopRequest string
	LogFormat string      // 日志格式
	TruncateDatabase bool // 是否进行建表导入数据操作
}


func ReadLine(filePth string,db *sql.DB,ranged_key []string) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	var request_index int
	var time_index int
	for i,v := range ranged_key {
		if v == "request" {
			request_index = i
		} else if v == "time_local" {
			time_index = i
		}

	}
	for {
		// 按行读取文件
		line, _ := bfRd.ReadBytes('\n')
		// 空行判断
		if len(line) != 0 {
			// byte转string
			str := string(line[:])
			// string split  类似于awk
			a := strings.Split(str, " |")
			// 将请求参数去除
			if strings.Contains(a[request_index],"?") {
				a[request_index] = strings.Split(a[request_index], "?")[0]
			}
			// 截取时间到小时
			time_local := strings.Split(a[time_index],":")
			a[time_index] = time_local[0] + "-" + time_local[1]
			InsertData(db,ranged_key,a)

		} else {
			return nil
		}
	}
	return nil
}

// 读入数据
func FilterTime(start string,end string, file_path string) {
	// build cmd
	var command string
	if len(end) != 0 {
		command = "sed -n '" + start + "," + end + "p' " + file_path + " > " + tmp_path
	} else {
		command = "sed -n '" + start + "p' " + file_path + " > " + tmp_path
	}
	// 使用bash才能够使用重定向 >
	c := exec.Command("bash","-c",command)
	//fmt.Println(c.Args)
	// debug cmd
	var stderr bytes.Buffer
	c.Stderr = &stderr
	err := c.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
}


// 读取日志格式
func ParseFormat(format string) []string {
	// 创建map
	str := strings.Split(format, "|")
	ranged_key := make([]string,0)
	for _,value := range str {
		// 删除prefix
		a := strings.TrimPrefix(value,"$")
		ranged_key = append(ranged_key,a)
	}
	return ranged_key
}

//todo
func ClearTmpFile()  {

}

//todo
func CheckArg()  {

}

//todo
func CheckConfig()  {

}
