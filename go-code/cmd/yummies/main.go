package main

import (
	"net/http"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/joseluis8906/yummies/go-code/internal/app"
)

func main() {
	fx.New(
		fx.Options(app.Module),
		fx.Provide(app.NewGRPCServer),
		fx.Provide(app.NewHTTPServer),

		fx.Invoke(func(*grpc.Server) {}),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
