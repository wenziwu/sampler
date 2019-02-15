package main

import (
	"github.com/sqshq/sampler/config"
	"github.com/sqshq/sampler/console"
	"github.com/sqshq/sampler/data"
	"github.com/sqshq/sampler/event"
	"github.com/sqshq/sampler/widgets"
	"github.com/sqshq/sampler/widgets/runchart"
	ui "github.com/sqshq/termui"
	"time"
)

func main() {

	cfg := config.Load()
	csl := console.Console{}
	csl.Init()
	defer csl.Close()

	width, height := ui.TerminalDimensions()
	layout := widgets.NewLayout(width, height, widgets.NewMenu())

	for _, c := range cfg.RunCharts {

		legend := runchart.Legend{Enabled: c.Legend.Enabled, Details: c.Legend.Details}
		chart := runchart.NewRunChart(c.Title, c.Precision, c.RefreshRateMs, legend)
		layout.AddComponent(chart, c.Title, c.Position, c.Size, widgets.TypeRunChart)

		for _, item := range c.Items {
			data.NewSampler(chart, item, c.RefreshRateMs)
		}
	}

	handler := event.Handler{
		Layout:        layout,
		RenderEvents:  time.NewTicker(console.RenderRate).C,
		ConsoleEvents: ui.PollEvents(),
	}

	handler.HandleEvents()
}
