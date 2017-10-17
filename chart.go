package main

import (
	"github.com/wcharczuk/go-chart"
	"bytes"
	"strconv"
	"encoding/base64"
)

func InitGraph()  {
	var src []chart.Value
	var s chart.Value
	for _,v := range data.TimeNumber {
		value,_ := strconv.ParseFloat(v[0],6)
		s.Value = value
		s.Label = v[1]
		src = append(src,s)
	}
	sbc := chart.BarChart{
		//定义图片样式
		Height:   512,
		XAxis: chart.Style{
			Show: true,
			FontSize: 8,
			TextRotationDegrees: 75.0,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		//读取数据
		Bars: src,
	}
	//生成图片
	buffer := bytes.NewBuffer([]byte{})
	err := sbc.Render(chart.PNG, buffer)
	Check(err)
	//直接以base64编码插入html
	images := base64.StdEncoding.EncodeToString(buffer.Bytes())
	data.Image = images

}