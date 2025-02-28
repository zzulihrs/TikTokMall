package main

import (
	"context"
	"github.com/tiktokmall/backend/app/order/biz/consumer"
	"github.com/tiktokmall/backend/app/order/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/order/biz/model"
	"github.com/tiktokmall/backend/app/order/infra/mq"
	"github.com/tiktokmall/backend/app/order/utils"
	"log"
	"net"
	"time"

	"github.com/tiktokmall/backend/app/order/biz/dal"
	"github.com/tiktokmall/backend/app/order/conf"
	"github.com/tiktokmall/backend/common/mtl"
	"github.com/tiktokmall/backend/common/serversuite"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/order/orderservice"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegisterAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	_ = godotenv.Load()

	mtl.InitMetric(ServiceName,
		conf.GetConf().Kitex.MetricsPort,
		RegisterAddr)
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background())
	dal.Init()
	mq.Init()
	consumer.Init()

	opts := kitexInit()

	svr := orderservice.NewServer(new(OrderServiceImpl), opts...)

	CornUpdateAllOrderStatus()

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}

}

/*
CornUpdateAllOrderStatus 创建一个corn任务，每天2点执行一次，更新超时订单状态
*/
func CornUpdateAllOrderStatus() {
	cm := utils.NewCronManager()

	defer cm.Stop()

	// 用于跟踪任务执行次数
	var counter int32

	jobID := "CornUpdateAllOrderStatus"
	_, err := cm.AddJob(jobID, "0 0 2 * * *", func() {
		order, err := model.GetAllOrder(mysql.DB, context.Background())
		if err != nil {
			log.Println("获取订单失败:", err)
			return
		}
		for _, v := range order {
			status, err := model.GetOrderStatus(mysql.DB, context.Background(), v.OrderId)
			if err != nil {
				log.Println("获取订单状态失败: ", err)
				return
			}
			// 判断v.CreatedAt超过1小时
			if status == 0 && time.Since(v.CreatedAt) > time.Hour {
				if err := model.ChangeOrderStatus(mysql.DB, context.Background(), v.OrderId, 3); err != nil {
					log.Println("更新订单状态失败:", err)
					return
				}
			}
		}
		log.Println("任务执行次数:", counter)

		return
	})

	if err != nil {
		log.Println("添加任务失败: %v", err)
	}

}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts,
		server.WithServiceAddr(addr),
		server.WithSuite(serversuite.CommonServerSuite{
			CurrentServiceName: ServiceName,
			RegisterAddr:       RegisterAddr,
		}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
