package main

import (
	"easy-chat/pkg/configserver"
	"easy-chat/pkg/resultx"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"sync"

	"easy-chat/apps/user/api/internal/config"
	"easy-chat/apps/user/api/internal/handler"
	"easy-chat/apps/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/dev/user.yaml", "the config file")
var wg sync.WaitGroup

//	func main() {
//		flag.Parse()
//
//		var c config.Config
//		err := configserver.NewConfigserver(*configFile, configserver.NewSail(&configserver.Config{
//			ETCDEndpoints:  "192.168.117.80:3379",
//			ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
//			Namespace:      "user",
//			Configs:        "user-api.yaml",
//			ConfigFilePath: "./etc/conf",
//			LogLevel:       "DEBUG",
//		})).MustLoad(&c, func(bytes []byte) error {
//			var c config.Config
//			configserver.LoadFromJsonBytes(bytes, &c)
//
//			proc.WrapUp()
//			wg.Add(1)
//			go func(c config.Config) {
//				defer wg.Done()
//				Run(c)
//			}(c)
//			return nil
//		})
//
//		if err != nil {
//			panic(err)
//		}
//		wg.Add(1)
//		go func(c config.Config) {
//			defer wg.Done()
//			Run(c)
//		}(c)
//
//		wg.Wait()
//	}
//
//	func Run(c config.Config) {
//		server := rest.MustNewServer(c.RestConf)
//		defer server.Stop()
//
//		ctx := svc.NewServiceContext(c)
//		handler.RegisterHandlers(server, ctx)
//
//		httpx.SetErrorHandlerCtx(resultx.ErrHandler(c.Name))
//		httpx.SetOkHandler(resultx.OkHandler)
//
//		fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
//		server.Start()
//	}
func main() {
	flag.Parse()

	var c config.Config
	//conf.MustLoad(*configFile, &c)

	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "192.168.88.128:3379",
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "user",
		Configs:        "user-api.yaml",
		ConfigFilePath: "./etc/dev",
		LogLevel:       "DEBUG",
	})).MustLoad(&c, func(bytes []byte) error {
		var c config.Config
		configserver.LoadFromJsonBytes(bytes, &c)

		proc.WrapUp()

		wg.Add(1)
		go func(c config.Config) {
			defer wg.Done()

			Run(c)
		}(c)
		return nil
	})
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func(c config.Config) {
		defer wg.Done()

		Run(c)
	}(c)

	wg.Wait()
}

func Run(c config.Config) {
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(resultx.ErrHandler(c.Name))
	httpx.SetOkHandler(resultx.OkHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
