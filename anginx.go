package main

import (
	"os"
	"bufio"
	"flag"
	"strings"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os/exec"
	"bytes"
	"database/sql"
	"log"
)

// 配置文件结构体
type Conf struct {
	Input_file string      // 需要分析的日志文件路径
	Start_date string      // 开始日期
	End_date string        // 结束日期
	Overtime float64       // 超时时间
	Log_format string      // 日志格式
}

var t Conf
var tmp_path string = "/tmp/tmp.log"
var result_path string = "Anginx.html"

func ReadLine(filePth string,db *sql.DB,ranged_key []string) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	var request_index int
	for i,v := range ranged_key {
		if v == "request" {
			request_index = i
		}
	}
	for {
		//按行读取文件
		line, _ := bfRd.ReadBytes('\n')
		//空行判断
		if len(line) != 0 {
			// byte转string
			str := string(line[:])
			// string split  类似于awk
			a := strings.Split(str, "|")
			//将请求参数去除
			if strings.Contains(a[request_index],"?") {
				a[request_index] = strings.Split(a[request_index], "?")[0]
			}
			InsertData(db,ranged_key,a)

		} else {
			return nil
		}
	}
	return nil
}

func FilterTime(start string,end string, file_path string) {
	//build cmd
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

func ParseFormat(format string) []string {
	//创建map
	str := strings.Split(format, "|")
	ranged_key := make([]string,0)
	for _,value := range str {
		//删除prefix
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

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main()  {
	config_path := flag.String("c","","please offer config")
	// 解析命令行参数
	flag.Parse()
	// 读取配置文件
	config,_:= ioutil.ReadFile(*config_path)
	//fmt.Println(string(config))
	t = Conf{}
	// 解析yaml文件
	yaml.Unmarshal(config, &t)
	fmt.Println("解析文件：",t.Input_file,"输出文件",result_path)
	//按日期过滤
	FilterTime(t.Start_date,t.End_date,t.Input_file)

	////初始化数据库
	//db := InitDatabase()
	////读取日志格式
	//ranged_key := ParseFormat(t.Log_format)
	////建表
	//CreateTable(db,ranged_key)
	////读入数据
	//ReadLine(tmp_path,db,ranged_key)
	////断开数据库连接
	//defer db.Close()
	//
	////统计独立ip数
	//CountUniqueIP(db)
	////统计请求数
	//CountRequest(db)
	////统计访问量前200的请求
	//ListPopularURL(db)

	InitTemplate()
}