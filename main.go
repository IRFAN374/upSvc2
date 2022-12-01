package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	// "github.com/IRFAN374/upSvc2/common/chttp"
	"github.com/IRFAN374/upSvc2/common/config"
	"github.com/gorilla/mux"

	"github.com/oklog/oklog/pkg/group"
	
	gokitlogzap "github.com/go-kit/kit/log/zap"
	// kitHttp "github.com/go-kit/kit/transport/http"
	gokitlog "github.com/go-kit/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var env string

func init() {
	flag.StringVar(&env, "env", "", "kube env")
}

func main() {

	fmt.Println("....... Welcome to golang jwt service.........")

	flag.Parse()

	if env == "" {
		os.Getenv("env")
	}

	cfg, err := config.NewConfig(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s-%s", cfg.ServiceName, env)

	// debugLogger, _, _, _ := getLogger(cfg.ServiceName, zapcore.DebugLevel)

	// var httpServerBefore = []kitHttp.ServerOption{
	// 	kitHttp.ServerErrorEncoder(kitHttp.ErrorEncoder(chttp.EncodeError)),
	// }
	// Htpp Middleware
	var mwf []mux.MiddlewareFunc

	// router
	httpRouter := mux.NewRouter().StrictSlash(false)
	httpRouter.Use(mwf...)
	httpRouter.PathPrefix("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

	})

	var server group.Group
	{
		httpServer := &http.Server{
			Addr:    ": " + strconv.Itoa(cfg.HttpPort),
			Handler: httpRouter,
		}

		server.Add(func() error {
			fmt.Printf("starting http server, port:%d \n", cfg.HttpPort)
			return httpServer.ListenAndServe()
		}, func(err error) {

			// write code here for gracefull shutDown

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
			defer cancel()

			httpServer.Shutdown(ctx)
		})

	}

	{
		cancelInterrupt := make(chan struct{})

		server.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}

		}, func(err error) {
			close(cancelInterrupt)
		})
	}

	fmt.Printf("exiting...  errors: %v\n", server.Run())
}


func getLogger(serviceName string, level zapcore.Level) (debugL, infoL, errorL gokitlog.Logger, zapLogger *zap.Logger) {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	ec.EncodeDuration = zapcore.StringDurationEncoder
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(ec)

	fw, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}

	core := zapcore.NewCore(encoder, fw, level)
	zapLogger = zap.New(core)

	debugL = gokitlogzap.NewZapSugarLogger(zapLogger, zap.DebugLevel)
	debugL = gokitlog.With(debugL, "service", serviceName)

	infoL = gokitlogzap.NewZapSugarLogger(zapLogger, zap.InfoLevel)
	infoL = gokitlog.With(infoL, "service", serviceName)

	errorL = gokitlogzap.NewZapSugarLogger(zap.New(zapcore.NewCore(encoder, os.Stderr, level)), zap.ErrorLevel)
	errorL = gokitlog.With(errorL, "service", serviceName)

	return

}
