package main


import (
	"html/template"
	"os"
)

type Data struct{
	UniqueIPNumber string
	RequestNumber string
	PopularURL [200][2]string
	PopularIP  [50][2]string
}


func GenerateHtml()  {
	// 声明模板内容
	const tpl = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Anginx</title>
    </head>
    <body>
		<div>
			<p>Unique IP Number : {{.UniqueIPNumber}}</p>
			<ul>
				{{ range $_,$value := .PopularIP }}
						<li>
							{{ range $value }}
								<span>{{.}}</span>
							{{ end }}
						</li>
				{{ end }}
			</ul>
		</div>
		<div>
			<p>Total Request Number : {{.RequestNumber}}<p>
				<ul>
				{{ range $_,$value := .PopularURL }}
						<li>
							{{ range $value }}
								<span>{{.}}</span>
							{{ end }}
						</li>
				{{ end }}
			</ul>
		</div>
    </body>
</html>`

	// 创建一个新的模板，并且载入内容
	t, err := template.New("Anginx.html").Parse(tpl)
	Check(err)
	// 定义传入到模板的数据，并生成html文件
	f, err := os.OpenFile(result_path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	Check(err)
	err = t.Execute(f, data)
	Check(err)
	defer f.Close()

}