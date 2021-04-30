/***************************
@File        : server.go
@Time        : 2021/04/19 13:48:11
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 华开教育小程序后台
****************************/

package main

import (
	"context"
	"fmt"
	"huakai/config"
	"huakai/handler"
	"huakai/pagination"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// 日志

func Logger() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//设置输出
	log.Out = src
	gin.DefaultWriter = src
	//设置日志级别
	log.SetLevel(logrus.DebugLevel)

	//设置日志格式
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return log
}

// 设置日志

func LoggerToFile() gin.HandlerFunc {
	logger := Logger()
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

// 关闭服务
func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	fmt.Println("got a signal", sig)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		fmt.Println("err", err)
	}
	// 看看实际退出所耗费的时间
	fmt.Println("------exited--------", time.Since(now))
}

// 跨域

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(LoggerToFile())
	// 日志文件不需要颜色
	// gin.DisableConsoleColor()
	r.Use(Cors())

	r.Static("/static", "./static")                                  // 设置静态路径
	paginationHandler := pagination.New(&pagination.CustomeParser{}) // 分页中间件
	log.WithFields(logrus.Fields{}).Info("Server Start ...")
	api := r.Group("/api")
	api.POST("/active", handler.CreateActive)                //新建活动
	api.GET("/active", paginationHandler, handler.GetActive) // 获取活动
	api.POST("/change/active", handler.UpdateActive)         // 修改活动
	api.POST("/sign", handler.CreateSign)                    // 申请活动
	api.GET("/sign", paginationHandler, handler.GetSign)     // 获取申请者
	api.PUT("/sign", handler.ChangeSign)
	api.POST("/join", handler.CreateJoinTable)
	api.GET("/join", paginationHandler, handler.GetJoinMes)
	api.PUT("/join", handler.ChangeJoinTable)
	api.POST("/member", handler.CreateMemberRes)
	api.GET("/member", paginationHandler, handler.GetMemberRes)
	api.POST("/change/member", handler.ChangeMemberRes)
	api.GET("/memberUser", paginationHandler, handler.GetMembers)
	api.DELETE("/delete", handler.DeleteData)
	api.GET("/imgList", paginationHandler, handler.GetImg)
	api.POST("/createImg", handler.CreateImage)
	api.POST("/changeImg", handler.ChangeImage)
	v1 := api.Group("/wx")
	{
		v1.POST("/user/login", handler.Login)
	}

	v2 := api.Group("/web")
	{
		v2.POST("/user/login", handler.LoginWeb)
		v2.GET("/getUsers", paginationHandler, handler.GetUsers)
		v2.POST("/addSuper", handler.AddSuperUser)
		v2.PUT("/changeUser", handler.ChangeUser)
	}

	// 首页
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/admin/login.html")
	})

	// 文档页面
	r.GET("/doc", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/doc/index.html")
	})
	addr := ":" + config.Con.Port
	server := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go server.ListenAndServe()
	gracefulExitWeb(server)
}
