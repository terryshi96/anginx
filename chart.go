package main

import (
	"github.com/wcharczuk/go-chart"
	"bytes"
	"os"
	"strconv"
	"fmt"
)

func InitGraph()  {
	var src []chart.Value
	var s chart.Value
	for _,v := range data.TimeNumber {
		value,_ := strconv.ParseFloat(v[0],6)
		s.Value = value
		s.Label = v[1]
		fmt.Println(s.Label)
		src = append(src,s)
	}
	sbc := chart.BarChart{
		Height:   512,
		BarWidth: 200,
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: src,
	}

	buffer := bytes.NewBuffer([]byte{})
	err := sbc.Render(chart.PNG, buffer)
	f,err := os.OpenFile("statics.png", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	f.Write(buffer.Bytes())
	Check(err)
}