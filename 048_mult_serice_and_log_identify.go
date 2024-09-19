package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// 管理并发任务
var g errgroup.Group

// 定义日志格式化结构体，本质还是一段字符串
type Log_formatter_with_instance_identify struct {
	instance_name string
}

// 日志格式化方法，最后构建为一个字符串
// 这里要进行带有服务具体实例的日志构造{多态的服务用例， 日志}
func (p *Log_formatter_with_instance_identify) default_Log_formatter_with_instance_identify(param gin.LogFormatterParams) string {
	var status_color, memthod_color, reset_color string
	// 根据配置文件，判断是否需要彩色输出
	if param.IsOutputColor() {
		status_color = param.StatusCodeColor()
		memthod_color = param.MethodColor()
		reset_color = param.ResetColor()
	}
	if param.Latency > time.Minute {
		// 如果请求时间超过1分钟，将其格式化为分钟:秒
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("[gin] Service:%s |%v |%s %3d %s | %13v | %15 |%s %-7s %s %s %s\n %s",
		p.instance_name,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		status_color, param.StatusCode, reset_color,
		param.Latency,
		param.ClientIP,
		memthod_color, param.Method, reset_color,
		param.Path,
		param.ErrorMessage,
	)
}

// gin.defaultLogFormatter() 是 private 叫不出來
//
//	func (p *LogFormatterWithInstanceIdentify) defaultLogFormatterWithInstanceIdentify2(param gin.LogFormatterParams) string {
//		outString := gin.defaultLogFormatter(param)
//		strings.Replace(outString, "[GIN]", fmt.Sprintf("[GIN] Service:%s |", p.instanceName), 1)
//		return outString
//	}
func main() {
	g.Go(func() error {
		router1 := gin.New()
		// 创建一个实例不做出任何初始化操作
		router1.Use(gin.Recovery())
		// 恢复返回一个从任何恐慌中恢复的中间件，并写入 500（如果有）。

		logFormatter := Log_formatter_with_instance_identify{instance_name: "01"}
		// 日志格式化方法调用，初始化结构服务为01
		router1.Use(gin.LoggerWithFormatter(logFormatter.default_Log_formatter_with_instance_identify))
		// 日志格式化方法调用，初始化结构服务为01
		router1.GET("/", Index1)
		// 路由调用
		return router1.Run(":12138")
	})

	g.Go(func() error {
		router2 := gin.New()
		router2.Use(gin.Recovery())
		logFormatter := Log_formatter_with_instance_identify{instance_name: "02"}
		router2.Use(gin.LoggerWithFormatter(logFormatter.default_Log_formatter_with_instance_identify))
		router2.GET("/", Index2)
		return router2.Run(":12139")
	})

	g.Wait()
}

func Index1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "hello server 1",
	})
}

func Index2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "hello server 2",
	})
}
