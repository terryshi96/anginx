package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

// 变量声明
var t Conf
var data Data
var tmp_path string = "/tmp/tmp.log"
var result_path string = "Anginx_9-30.html"


// 抛出异常
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
	t = Conf{}
	// 解析yaml文件
	yaml.Unmarshal(config, &t)
	fmt.Println("解析文件：",t.InputFile,"输出文件",result_path)
	// 按日期过滤
	FilterTime(t.StartDate,t.EndDate,t.InputFile)

	// 初始化数据库
	db := InitDatabase()
	if t.TruncateDatabase {
		ranged_key := ParseFormat(t.LogFormat)
		CreateTable(db, ranged_key)
		ReadLine(tmp_path, db, ranged_key)
	}
	// 断开数据库连接
	defer db.Close()

	// 数据分析
	data.UniqueIPNumber = CountUniqueIP(db)
	data.RequestNumber = CountRequest(db)
	data.PopularURL = ListPopularURL(db)
	data.PopularIP = ListPopularIP(db)
	data.OvertimeReq = ListAvg(db)
	data.LongestReq = ListLongest(db)
	data.ErrorReq,data.ErrorRate = ListError(db)
	//InitGraph()
	// 生成html
	GenerateHtml()

}