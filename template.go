package main


import (
	"html/template"
	"os"
	"bytes"
)

type Data struct{
	UniqueIPNumber string
	RequestNumber string
	PopularURL [][2]string
	PopularIP  [][2]string
	OvertimeReq [][2]string
	LongestReq [][2]string
	ErrorRate  string
	ErrorReq   [][2]string
	Chart      *bytes.Buffer
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
			<p>Overtime Requests</p>
			<ul>
				<span>Avg Request Time</span><span>Request Detail</span>
				{{ range $_,$value := .OvertimeReq }}
						<li>
							{{ range $value }}
								<span>{{.}}</span>
							{{ end }}
						</li>
				{{ end }}
			</ul>
			<ul>
				<span>Longest Request Times	</span><span>Request Detail</span>
				{{ range $_,$value := .LongestReq }}
						<li>
							{{ range $value }}
								<span>{{.}}</span>
							{{ end }}
						</li>
				{{ end }}
			</ul>
		</div>
		<div>
			<p>Error Rate : {{.ErrorRate}}%<p>
			<p>Error Requests</p>
			<ul>
				<span>Error Code</span><span>Request Detail</span>
				{{ range $_,$value := .ErrorReq }}
						<li>
							{{ range $value }}
								<span>{{.}}</span>
							{{ end }}
						</li>
				{{ end }}
			</ul>
		</div>
		<div>
			<p>Unique IP Number : {{.UniqueIPNumber}}</p>
			<p>IP Statistics</p>
			<ul>
				<span>Request Times	</span><span>IP</span>
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
			<p>Requests Statistics</p>
			<ul>
				<span>Request Times</span><span>Request Detail</span>
				{{ range $_,$value := .PopularURL }}
						<li>
							{{ range $value }}
								<span>{{.}}</span>
							{{ end }}
						</li>
				{{ end }}
			</ul>
		</div>
		<div>
		   <img>src="{{ .Chart }}"</img>
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