package charts

import "github.com/go-echarts/go-echarts/types"

// Scatter represents a scatter chart.
type Scatter struct {
	RectChart
}

func (Scatter) Type() string { return types.ChartScatter }

// NewScatter creates a new scatter chart.
func NewScatter() *Scatter {
	chart := &Scatter{}
	chart.initBaseConfiguration()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *Scatter) SetXAxis(xAxis interface{}) *Scatter {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *Scatter) AddSeries(name string, yAxis interface{}, opts ...SeriesOpts) *Scatter {
	series := SingleSeries{Name: name, Type: types.ChartScatter, Data: yAxis}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
