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

	"github.com/IRFAN374/upSvc2/common/chttp"
	"github.com/IRFAN374/upSvc2/common/config"
	"github.com/IRFAN374/upSvc2/db"
	token "github.com/IRFAN374/upSvc2/token"
	tokenSvcMw "github.com/IRFAN374/upSvc2/token/service"

	user "github.com/IRFAN374/upSvc2/user"
	userMw "github.com/IRFAN374/upSvc2/user/service"
	usertransport "github.com/IRFAN374/upSvc2/user/transport"
	usertransporthttp "github.com/IRFAN374/upSvc2/user/transport/http"

	tokenRepository "github.com/IRFAN374/upSvc2/reposiotry/token"
	tokenRedisRepo "github.com/IRFAN374/upSvc2/reposiotry/token/redis"
	tokenMw "github.com/IRFAN374/upSvc2/reposiotry/token/service"

	userRepository "github.com/IRFAN374/upSvc2/reposiotry/user"
	userMongo "github.com/IRFAN374/upSvc2/reposiotry/user/mongo"
	userRepoMw "github.com/IRFAN374/upSvc2/reposiotry/user/service"

	redis "github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"

	"github.com/oklog/oklog/pkg/group"

	gokitlogzap "github.com/go-kit/kit/log/zap"
	kitHttp "github.com/go-kit/kit/transport/http"
	log "github.com/go-kit/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var env string

func init() {
	flag.StringVar(&env, "env", "", "kube env")
}

var redisClient *redis.Client

func init() {
	//Initializing redis
	redisClient = db.ConnectRedis()
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

	debugLogger, _, _, _ := getLogger(cfg.ServiceName, zapcore.DebugLevel)

	var httpServerBefore = []kitHttp.ServerOption{
		kitHttp.ServerErrorEncoder(kitHttp.ErrorEncoder(chttp.EncodeError)),
	}
	// Htpp Middleware
	var mwf []mux.MiddlewareFunc

	// router
	httpRouter := mux.NewRouter().StrictSlash(false)
	httpRouter.Use(mwf...)
	httpRouter.PathPrefix("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

	})

	// token repo is added

	var tokenRepo tokenRepository.Repository
	{
		tokenRepo = tokenRedisRepo.NewRepository(redisClient)
		tokenRepo = tokenMw.LoggingMiddleware(log.With(debugLogger, "repository", "token repository"))(tokenRepo)
	}

	// token server is added

	var tokenSvc token.Service
	{
		tokenSvc = token.NewService(tokenRepo)
		tokenSvc = tokenSvcMw.LoggingMiddleware(log.With(debugLogger, "service", "token service"))(tokenSvc)
	}

	// user repo is added

	var userRepo userRepository.Repository
	{
		userRepo = userMongo.NewMongoReposiotry()
		userRepo = userRepoMw.LoggingMiddleware(log.With(debugLogger, "reposiotry", "user reposiotry"))(userRepo)
	}

	// user server is added

	var userSvc user.Service
	{
		userSvc = user.NewService(tokenSvc, userRepo)
		userSvc = userMw.LoggingMiddleware(log.With(debugLogger, "service", "user service"))(userSvc)
		userEndpoints := usertransport.Endpoints(userSvc)
		userSvcHandler := usertransporthttp.NewHTTPHandler(&userEndpoints, httpServerBefore...)
		httpRouter.PathPrefix("/api").Handler(userSvcHandler)
	}

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

func getLogger(serviceName string, level zapcore.Level) (debugL, infoL, errorL log.Logger, zapLogger *zap.Logger) {
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
	debugL = log.With(debugL, "service", serviceName)

	infoL = gokitlogzap.NewZapSugarLogger(zapLogger, zap.InfoLevel)
	infoL = log.With(infoL, "service", serviceName)

	errorL = gokitlogzap.NewZapSugarLogger(zap.New(zapcore.NewCore(encoder, os.Stderr, level)), zap.ErrorLevel)
	errorL = log.With(errorL, "service", serviceName)

	return

}
