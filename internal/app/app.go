package app

import (
	"cms-service/pkg/pgxdb"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type App struct {
	deviceServiceProvider  *deviceServiceProvider
	commandServiceProvider *commandServiceProvider
	geozoneServiceProvider *geozoneServiceProvider
	ginServer              *gin.Engine
	db                     *pgxpool.Pool
}

func NewApp(ctx context.Context) (*App, error) {
	a := App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *App) Run() error {
	return a.ginServer.Run(":8000")
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initDb,
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

func (a *App) initDb(ctx context.Context) error {
	client, err := pgxdb.NewClient(ctx, "localhost", "1234",
		"postgres", "lux12345", "postgres")
	if err != nil {
		return err
	}

	a.db = client
	return nil
}

func (a *App) initDeviceServiceProvider(ctx context.Context) error {
	a.deviceServiceProvider = newDeviceServiceProvider(a.db)
	a.commandServiceProvider = newCommandServiceProvider(a.db)
	a.geozoneServiceProvider = newGeozoneServiceProvider(a.db)
	return nil
}

func (a *App) initGinServer(ctx context.Context) error {
	engine := gin.New()
	a.ginServer = engine
	log.Println(engine)
	a.deviceServiceProvider.DeviceImpl(engine)
	a.commandServiceProvider.CommandImpl(engine)
	a.geozoneServiceProvider.GeozoneImpl(engine)
	return nil
}
