package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
)

// Gauge represents a gauge chart.
type Gauge struct {
	BaseOpts
	Series
}

func (Gauge) chartType() string { return common.ChartType.Gauge }

// NewGauge creates a new gauge chart.
func NewGauge(routers ...RouterOpts) *Gauge {
	chart := new(Gauge)
	chart.initBaseOpts(routers...)
	return chart
}

// Add adds new data sets.
func (c *Gauge) Add(name string, data map[string]interface{}, options ...seriesOptser) *Gauge {
	nvs := make([]common.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, common.NameValueItem{Name: k, Value: v})
	}
	series := singleSeries{Name: name, Type: common.ChartType.Gauge, Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	return c
}

// SetGlobalOptions sets options for the Gauge instance.
func (c *Gauge) SetGlobalOptions(options ...globalOptser) *Gauge {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Gauge) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Gauge) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
