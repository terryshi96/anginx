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
)

// 配置文件结构体
type Conf struct {
	Input_file string
	Output_file string
	Start_date string
	End_date string
	Overtime float64
	Log_format string
}

var t Conf

func ReadLine(filePth string,db *sql.DB,ranged_key []string) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	for {
		//按行读取文件
		line, _ := bfRd.ReadBytes('\n')
		//空行判断
		if len(line) != 0 {
			// byte转string
			str := string(line[:])
			// string split  类似于awk
			a := strings.Split(str, "|")
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
		command = "sed -n '" + start + "," + end + "p' " + file_path + " > tmp"
	} else {
		command = "sed -n '" + start + "p' " + file_path + " > tmp"
	}
	// 使用bash才能够使用重定向 >
	c := exec.Command("bash","-c",command)
	fmt.Println(c.Args)
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
	fmt.Println("解析文件：",t.Input_file,"输出文件",t.Output_file)
	//按日期过滤
	FilterTime(t.Start_date,t.End_date,t.Input_file)
	//读取日志格式
	ranged_key := ParseFormat(t.Log_format)

	//初始化数据库
	db := InitDatabase()
	//建表
	CreateTable(db,ranged_key)
	//读入数据
	ReadLine("tmp",db,ranged_key)
	defer db.Close()
}