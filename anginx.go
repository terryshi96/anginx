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
)

// 配置文件结构体
type Conf struct {
	Input_file string
	Output_file string
	Start_date string
	End_date string
	Overtime float64
}

var t Conf

//hookfunc 对line进行处理
//func processLine(line []byte) {
//	//os.Stdout.Write(line)
//}
//func ReadLine(filePth string, hookfn func([]byte)) error {
func ReadLine(filePth string,result_path string) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)

	result_f, err := os.OpenFile(result_path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY|os.O_APPEND, 0666)

	for {
		//按行读取文件
		line, err := bfRd.ReadBytes('\n')
		//hookfn(line) //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		// byte转string
		str := string(line[:])
		// string split
		a := strings.Split(str," ")
		// string转float
		cost,_ := strconv.ParseFloat(a[4],3)
		if cost > t.Overtime {
			result_f.Write(line)
		}
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func Sed(filePath string) {

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
	ReadLine(t.Input_file,t.Output_file)
}