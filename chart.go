package main

import (
	"github.com/wcharczuk/go-chart"
	"bytes"
	"fmt"
)

func InitGraph()  {
	sbc := chart.BarChart{
		Height:   512,
		BarWidth: 60,
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: []chart.Value{
			{Value: 5.25, Label: "Blue"},
			{Value: 4.88, Label: "Green"},
			{Value: 4.74, Label: "Gray"},
			{Value: 3.22, Label: "Orange"},
			{Value: 3, Label: "Test"},
			{Value: 2.27, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := sbc.Render(chart.PNG, buffer)
	data.Chart = buffer
	Check(err)
}