package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

type Log_formatter_with_instance_identify struct {
	instance_name string
}

func (p *Log_formatter_with_instance_identify) default_Log_formatter_with_instance_identify(param gin.LogFormatterParams) string {
	var status_color, memthod_color, reset_color string
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
		router1.Use(gin.Recovery())
		logFormatter := Log_formatter_with_instance_identify{instance_name: "01"}
		router1.Use(gin.LoggerWithFormatter(logFormatter.default_Log_formatter_with_instance_identify))
		router1.GET("/", Index1)
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
