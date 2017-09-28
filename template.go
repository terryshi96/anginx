package main


import (
	"html/template"
	"os"
)

var Data struct{

}

func InitTemplate()  {
	// 声明模板内容
	const tpl = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Anginx</title>
    </head>
    <body>
    </body>
</html>`



	// 创建一个新的模板，并且载入内容
	t, err := template.New("Anginx").Parse(tpl)
	Check(err)

	// 定义传入到模板的数据，并生成html文件
	f := os.Open(result_path)
	err = t.Execute(os., Data)
	Check(err)

}