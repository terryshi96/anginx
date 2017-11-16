package main

import (
	"os/exec"
	"bytes"
	"fmt"
	"strings"
	"database/sql"
	"os"
	"bufio"
	"flag"
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
    EmailConfig struct {
    	Sending bool
    	UserName string
    	Password string
    	SmtpHost string
    	Sender string
    	Receivers []string
	}
}

//检查配置参数
func CheckConfig(t Conf) {
	if len(t.InputFile) == 0 {
		fmt.Println("Input file can not be empty")
	} else if len(t.StartDate) == 0 {
		fmt.Println("Start date can not be empty")
	} else if len(t.Overtime) == 0 {
		t.Overtime = "1"
	} else if len(t.TopIP) == 0 {
		t.TopIP = "10"
	} else if len(t.TopRequest) == 0 {
		t.TopRequest = "10"
	} else if len(t.LogFormat) == 0 {
		fmt.Println("Log format can not be empty")
	}
}


//读取命令行参数
func CheckArg() *string {
	usage := `
	please use a config file including
	inputfile---log file to parse
	startdate---use nginx date format
	enddate---use nginx date format
	overtime---over this time will be recorded
	topip---top x ip
	toprequest---top x requests
	logformat---nginx log format and split by " |"
		remote_addr request request_time status time_local must be included
	truncatedatabase---reload data or not
	`
	config_path := flag.String("c","",usage)
	h := flag.Bool("h", false, "help")
	// 解析命令行参数
	flag.Parse()
	if len(os.Args) == 1 {
		*h = true
	}
	if *h  {
		flag.Usage()
		os.Exit(1)
	}
	return config_path
}


// 读入数据
func ReadLine(filePth string,db *sql.DB,ranged_key []string) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	var request_index int
	//var time_index int
	for i,v := range ranged_key {
		if v == "request" {
			request_index = i
		} else if v == "time_local" {
			//time_index = i
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
			a[request_index] = strings.Replace(a[request_index],"HTTP/1.1","", -1)
			// 截取时间到小时
			//time_local := strings.Split(a[time_index],":")
			//a[time_index] = time_local[0] + "/" + time_local[1]
			InsertData(db,ranged_key,a)

		} else {
			return nil
		}
	}
	return nil
}


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
	ranged_key := make([]string, 0)
	if strings.Contains(format," |") && has(format) {
		str := strings.Split(format, " |")
		for _, value := range str {
			// 删除prefix
			a := strings.TrimPrefix(value, "$")
			ranged_key = append(ranged_key, a)
		}
	} else {
		fmt.Println("log should be splited by ' |' or has key fields")
		os.Exit(1)
	}

	return ranged_key
}

//时间字符串处理
func FormatTime(start string,end string)  {
	if strings.Contains(start,"\\") {
		a := strings.Replace(start, "\\", "", -1)
		b := strings.TrimPrefix(a, "/")
		data.StartDate = strings.TrimSuffix(b, "/")
		a = strings.Replace(end, "\\", "", -1)
		b = strings.TrimPrefix(a, "/")
		data.EndDate = strings.TrimSuffix(a, "/")
	} else {
		fmt.Println("date should be like '/11\\/Oct\\/2017/'")
		os.Exit(1)
	}
}

//判断关键字段是否包含
func has(format string) bool {
	if strings.Contains(format,"remote_addr") && strings.Contains(format,"request_time") && strings.Contains(format,"request") && strings.Contains(format,"status") && strings.Contains(format,"time_local") {
		return true
	}
	return false
}



