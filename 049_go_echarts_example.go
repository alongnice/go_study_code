package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	// 地图相关第三方库
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// 日志中间件，记录所有进入http服务器的请求 传递给下一个函数
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//返回http.handlerfunc 重载后的内容
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		// 重写日志参数
		handler.ServeHTTP(w, r)

	})
}

// 地图数据，定义部分省份，生成随机数据
var (
	baseMapData = []opts.MapData{
		{"北京", float64(rand.Intn(150))},
		{"上海", float64(rand.Intn(150))},
		{"广东", float64(rand.Intn(150))},
		{"辽宁", float64(rand.Intn(150))},
		{"山东", float64(rand.Intn(150))},
		{"山西", float64(rand.Intn(150))},
		{"陕西", float64(rand.Intn(150))},
		{"新疆", float64(rand.Intn(150))},
		{"内蒙古", float64(rand.Intn(150))},
	}
	// 市级数据
	guangdongMapData = map[string]float64{
		"深圳市": float64(rand.Intn(150)),
		"广州市": float64(rand.Intn(150)),
		"湛江市": float64(rand.Intn(150)),
		"汕头市": float64(rand.Intn(150)),
		"东莞市": float64(rand.Intn(150)),
		"佛山市": float64(rand.Intn(150)),
		"云浮市": float64(rand.Intn(150)),
		"肇庆市": float64(rand.Intn(150)),
		"梅州市": float64(rand.Intn(150)),
	}
)

// 地图数据生成器，将map类型数据转为[]opts.MapData类型
func generate_map_data(data map[string]float64) (items []opts.MapData) {
	items = make([]opts.MapData, 0)
	for k, v := range data {
		items = append(items, opts.MapData{Name: k, Value: v})
	}
	return
}

func mapBase() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "map-example",
		}),
	)
	mc.AddSeries("map", baseMapData)
	return mc
}

func mapShowLabel() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "map-example",
		}),
	)

	mc.AddSeries("map", baseMapData).SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show: opts.Bool(true), // Fix: Explicit boolean type
		}),
	)
	return mc
}
func mapVisualMap() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "map-Visualmap",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Show: opts.Bool(true), // Fix: Explicit boolean type
		}),
	)
	mc.AddSeries("map", baseMapData)
	return mc
}
func mapRegion() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "map_religion_guang_dong",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Show: opts.Bool(true), // Fix: Explicit boolean type
			InRange: &opts.VisualMapInRange{
				Color: []string{"#50a3ba", "#eac736", "#d94e5d"}},
		}),
	)
	mc.AddSeries("map", generate_map_data(guangdongMapData))
	return mc
}

func mapTheme() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "macarons",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Map-theme",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Show: opts.Bool(true), // Fix: Explicit boolean type
			Max:  150,
		}),
	)
	mc.AddSeries("map", baseMapData)
	return mc
}

func main() {
	page := components.NewPage() //创建一个页面
	// 添加基础地图，显示标签等
	page.AddCharts(
		mapBase(),
		mapShowLabel(),
		mapVisualMap(),
		mapRegion(),
		mapTheme(),
	) //添加图表
	f, err := os.Create("map.html")
	// 创建出地图文件
	if err != nil {
		panic(err)
	}

	page.Render(io.MultiWriter(f)) //将页面渲染到文件中

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "map.html") // 处理根路径并返回 map.html 文件
	})

	fs := http.FileServer(http.Dir("./"))
	log.Println("running server at 12138")
	log.Fatal(http.ListenAndServe("127.0.0.1:12138", logRequest(fs)))
}
