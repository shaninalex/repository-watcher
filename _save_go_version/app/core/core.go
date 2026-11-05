package core

type App struct {
	//Commands *bus.CommandBus
	//Events   *bus.EventBus
}

type Module interface {
	Name() string
	Register(a *App) error
}

type Resolve func() *App
