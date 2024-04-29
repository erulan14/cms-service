package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	deviceServiceProvider *deviceServiceProvider
	ginServer             *gin.Engine
}

func NewApp(ctx context.Context) (App, error) {
	a := App{}
	return a, nil
}

func (a *App) Run() error {
	return a.ginServer.Run(":8000")
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initDeviceServiceProvider,
		a.initGinServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	return nil
}

func (a *App) initDeviceServiceProvider(ctx context.Context) error {
	a.deviceServiceProvider = newDeviceServiceProvider()
	//a.deviceServiceProvider.db =
	return nil
}

func (a *App) initGinServer(ctx context.Context) error {
	engine := gin.New()
	a.ginServer = engine
	log.Println(engine)
	a.deviceServiceProvider.DeviceImpl(engine)
	return nil
}
