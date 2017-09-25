package main

import (
	"os"
	"bufio"
	"io"
	"flag"
	"strings"
	"strconv"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os/exec"
	"bytes"
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

//hookfunc 对line进行处理
//func processLine(line []byte) {
//	//os.Stdout.Write(line)
//}
//func ReadLine(filePth string, hookfn func([]byte)) error {
func ReadLine(filePth string,result_path string,ranged_key []string) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)

	//打开文件模式 文件不存在则创建 已有文件则清空 末尾添加 只写
	result_f, err := os.OpenFile(result_path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY|os.O_APPEND, 0666)

	for {
		//按行读取文件
		line, err := bfRd.ReadBytes('\n')
		//hookfn(line) //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。

		//空行判断
		if len(line) != 0 {
			var m = map[string]string{}
			// byte转string
			str := string(line[:])
			// string split  类似于awk
			a := strings.Split(str, "|")
			index := 0
			//值匹配 遍历map时key是随机化的 不能直接遍历
			for _,key := range ranged_key {
				m[key] = a[index]
				index++
			}
			//fmt.Println(m)
			// string转float
			cost, _ := strconv.ParseFloat(m["request_time"], 3)
			if cost > t.Overtime {
				result_f.Write(line)
			}
			if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
				if err == io.EOF {
					return nil
				}
				return err
			}
		} else {
			return nil
		}
	}
	return nil
}

func FilterTime(start string,end string, file_path string) {
	//build cmd
	command := "sed -n '" + start + "," + end + "p' " + file_path + " > tmp"
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
	FilterTime(t.Start_date,t.End_date,t.Input_file)
	ranged_key := ParseFormat(t.Log_format)
	ReadLine("tmp",t.Output_file,ranged_key)
}