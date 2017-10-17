package main


import (
	"html/template"
	"os"
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
	TimeNumber [][2]string
	StartDate string
	EndDate string
	Image 	string
}


func GenerateHtml()  {
	// 声明模板内容
	const tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Nginx Analysis</title>
    <link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Nosifer" rel="stylesheet">
    <script src="https://cdn.bootcss.com/jquery/2.1.1/jquery.min.js"></script>
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
    <style>
        #title {
            font-family: 'Nosifer', cursive;
            text-align: center;
        }

        div.panel.panel-default {
            width: 300px;
            margin: auto;
        }

        div.common {
            width: 300px;
            margin-bottom: 0px;
            padding-left: 15px;
            padding-right: 15px;
            padding-top: 10px;
            padding-bottom: 10px;
            background-color: #f5f5f5;
            border: solid 1px #ddd;
            border-radius: 3px;
            margin: auto;
        }

        #image {
            text-align: center;
        }

        div.panel-body {
            width: 800px;
            margin: auto;
        }

        td {
            word-wrap: break-word;
        }
    </style>
</head>
<body>
<h1 id="title">
    Anginx
</h1>
<div class="common">
    <span>{{.StartDate}} {{.EndDate}}</span>
</div>

<div class="panel-group" id="accordion">

    <div class="panel panel-default">
        <div class="panel-heading">
            <h4 class="panel-title">
                <a data-toggle="collapse" data-parent="#accordion" href="#collapse1">
                    Overtime Requests(Avg)
                </a>
            </h4>
        </div>
    </div>
    <div id="collapse1" class="panel-collapse collapse">
        <div class="panel-body">
            <table class="table">
                <thead>
                <tr>
                    <th>Avg Request Time(s)</th>
                    <th>Request Detail</th>
                </tr>
                </thead>
                <tbody>
                {{ range $_,$value := .OvertimeReq }}
                <tr>
                    {{ range $value }}
                    <td>{{.}}</td>
                    {{ end }}
                </tr>
                {{ end }}
                <tbody>
            </table>
        </div>
    </div>

    <div class="panel panel-default">
        <div class="panel-heading">
            <h4 class="panel-title">
                <a data-toggle="collapse" data-parent="#accordion" href="#collapse2">
                    Overtime Requests(Longest)
                </a>
            </h4>
        </div>
    </div>
    <div id="collapse2" class="panel-collapse collapse">
        <div class="panel-body">
            <table class="table">
                <thead>
                <tr>
                    <th>Longest Request Time(s)</th>
                    <th>Request Detail</th>
                </tr>
                </thead>
                <tbody>
                {{ range $_,$value := .LongestReq }}
                <tr>
                    {{ range $value }}
                    <td>{{.}}</td>
                    {{ end }}
                </tr>
                {{ end }}
                <tbody>
            </table>
        </div>
    </div>


    <div class="common">
        <span>Error Rate : {{.ErrorRate}}%</span>
    </div>

    <div class="panel panel-default">
        <div class="panel-heading">
            <h4 class="panel-title">
                <a data-toggle="collapse" data-parent="#accordion" href="#collapse3">
                    Error Requests
                </a>
            </h4>
        </div>
    </div>
    <div id="collapse3" class="panel-collapse collapse">
        <div class="panel-body">
            <table class="table">
                <thead>
                <tr>
                    <th>Error Code</th>
                    <th>Request Detail</th>
                </tr>
                </thead>
                <tbody>
                {{ range $_,$value := .ErrorReq }}
                <tr>
                    {{ range $value }}
                    <td>{{.}}</td>
                    {{ end }}
                </tr>
                {{ end }}
                <tbody>
            </table>
        </div>
    </div>

    <div class="common">
        <span>Unique IP Number : {{.UniqueIPNumber}}</span>
    </div>

    <div class="panel panel-default">
        <div class="panel-heading">
            <h4 class="panel-title">
                <a data-toggle="collapse" data-parent="#accordion" href="#collapse4">
                    IP Number Statistics
                </a>
            </h4>
        </div>
    </div>
    <div id="collapse4" class="panel-collapse collapse">
        <div class="panel-body">
            <ul class="list-group">
                <table class="table">
                    <thead>
                    <tr>
                        <th>Request Number</th>
                        <th>IP</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{ range $_,$value := .PopularIP }}
                    <tr>
                        {{ range $value }}
                        <td>{{.}}</td>
                        {{ end }}
                    </tr>
                    {{ end }}
                    <tbody>
                </table>
        </div>
    </div>

    <div class="common">
        <span>Total Request Number : {{.RequestNumber}}</span>
    </div>
    <div class="panel panel-default">
        <div class="panel-heading">
            <h4 class="panel-title">
                <a data-toggle="collapse" data-parent="#accordion" href="#collapse5">
                    Requests Number Statistics
                </a>
            </h4>
        </div>
    </div>
    <div id="collapse5" class="panel-collapse collapse">
        <div class="panel-body">
            <table class="table">
                <thead>
                <tr>
                    <th>Request Number</th>
                    <th>Request Detail</th>
                </tr>
                </thead>
                <tbody>
                {{ range $_,$value := .PopularURL }}
                <tr>
                    {{ range $value }}
                    <td>{{.}}</td>
                    {{ end }}
                </tr>
                {{ end }}
                <tbody>
            </table>
        </div>
    </div>

    <div class="panel panel-default">
        <div class="panel-heading">
            <h4 class="panel-title">
                <a data-toggle="collapse" data-parent="#accordion" href="#collapse6">
                    Request Number Per Hour
                </a>
            </h4>
        </div>
    </div>
    <div id="collapse6" class="panel-collapse collapse">
        <div id="image" class="panel-body">
            <img src="data:image/png;base64,{{.Image}}" alt="img"/>
        </div>
    </div>
</div>
</body>
</html>
`

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