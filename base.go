package goecharts

// 图形初始化配置项
type InitOpts struct {
	// 生成的 HTML 页面标题
	PageTitle string `default:"Awesome go-echarts"`
	// 画布宽度
	Width string `default:"800px"`
	// 画布高度
	Height string `default:"500px"`
	// 图表 ID，是图表唯一标识
	ChartID string
	// 静态资源 host 地址
	AssetsHost string `default:"http://chenjiandongx.com/go-echarts-assets/assets/"`
	// 图表主题
	Theme string `default:"white"`
}

// 静态资源配置项
type AssetsOpts struct {
	JSAssets  orderSet
	CSSAssets orderSet
}

// 初始化静态资源配置项
func (opt *AssetsOpts) initAssetsOpts() {
	opt.JSAssets.init("echarts.min.js")
	opt.CSSAssets.init("bulma.min.css")
}

// 返回资源列表
func (opt *AssetsOpts) yieldAssets() ([]string, []string) {
	return opt.JSAssets.Values, opt.CSSAssets.Values
}

// 校验静态资源配置项，追加 host
func (opt *AssetsOpts) validateAssets(host string) {
	for i := 0; i < len(opt.JSAssets.Values); i++ {
		opt.JSAssets.Values[i] = host + opt.JSAssets.Values[i]
	}
	for j := 0; j < len(opt.CSSAssets.Values); j++ {
		opt.CSSAssets.Values[j] = host + opt.CSSAssets.Values[j]
	}
}

// 为 InitOptions 设置字段默认值
func (opt *InitOpts) setDefault() error {
	return setDefaultValue(opt)
}

// 确保 ContainerID 不为空且唯一
func (opt *InitOpts) checkID() {
	if opt.ChartID == "" {
		opt.ChartID = genChartID()
	}
}

// 验证初始化参数，确保图形能够得到正确渲染
func (opt *InitOpts) validateInitOpt() {
	opt.setDefault()
	opt.checkID()
}

// Http 路由
type HTTPRouter struct {
	URL  string // 路由 URL
	Text string // 路由显示文字
}

type HTTPRouters []HTTPRouter

func (hr HTTPRouters) Len() int {
	return len(hr)
}

// 全局颜色配置项
type ColorOpts []string

// 所有图表都拥有的基本配置项
type BaseOpts struct {
	InitOpts             // 图形初始化配置项
	LegendOpts           // 图例组件配置项
	TooltipOpts          // 提示框组件配置项
	ToolboxOpts          // 工具箱组件配置项
	TitleOpts            // 标题组件配置项
	AssetsOpts           // 静态资源配置项
	ColorList   []string // 全局颜色列表
	appendColor []string // 追加全局颜色列表
	HTTPRouters          // 路由列表
	DataZoomOptsList     // 区域缩放组件配置项列表
	VisualMapOptsList    // 视觉映射组件配置项列表
	GeoOpts              // 地理坐标系组件配置项

	HasXYAxis bool // 图形是否拥有 XY 轴
}

// 设置全局颜色
func (opt *BaseOpts) setColor(options ...interface{}) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case ColorOpts:
			opt.appendColor = append(opt.appendColor, option.(ColorOpts)...)
		}
	}
}

// 初始化全局颜色列表
func (opt *BaseOpts) initSeriesColors() {
	opt.ColorList = []string{
		"#c23531", "#2f4554", "#61a0a8", "#d48265", "#91c7ae", "#749f83",
		"#ca8622", "#bda29a", "#6e7074", "#546570", "#c4ccd3"}
}

// 初始化 BaseOpts
func (opt *BaseOpts) initBaseOpts(routers ...HTTPRouter) {
	for i := 0; i < len(routers); i++ {
		opt.HTTPRouters = append(opt.HTTPRouters, routers[i])
	}
	opt.initSeriesColors()
}

// 插入颜色到颜色列表首部
func (opt *BaseOpts) insertSeriesColors(s []string) {
	tmpCl := reverseSlice(s) // 翻转颜色列表
	// 颜色追加至首部
	for i := 0; i < len(tmpCl); i++ {
		opt.ColorList = append(opt.ColorList, "")
		copy(opt.ColorList[1:], opt.ColorList[0:])
		opt.ColorList[0] = tmpCl[i]
	}
}

// 设置 BaseOptions 全局配置项
func (opt *BaseOpts) setBaseGlobalConfig(options ...interface{}) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case InitOpts:
			opt.InitOpts = option.(InitOpts)
			if opt.InitOpts.Theme != "" {
				opt.JSAssets.Add("themes/" + opt.Theme + ".js")
			}
		case TitleOpts:
			opt.TitleOpts = option.(TitleOpts)
		case ToolboxOpts:
			opt.ToolboxOpts = option.(ToolboxOpts)
		case TooltipOpts:
			opt.TooltipOpts = option.(TooltipOpts)
		case LegendOpts:
			opt.LegendOpts = option.(LegendOpts)
		case ColorOpts:
			opt.insertSeriesColors(option.(ColorOpts))
		case DataZoomOpts:
			opt.DataZoomOptsList = append(opt.DataZoomOptsList, option.(DataZoomOpts))
		case VisualMapOpts:
			opt.VisualMapOptsList = append(opt.VisualMapOptsList, option.(VisualMapOpts))
		}
	}
}
