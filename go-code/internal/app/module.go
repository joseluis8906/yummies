package app

import (
	"go.uber.org/fx"

	"github.com/joseluis8906/yummies/go-code/internal/nats"
	"github.com/joseluis8906/yummies/go-code/internal/config"
	"github.com/joseluis8906/yummies/go-code/internal/mongodb"
	"github.com/joseluis8906/yummies/go-code/internal/log"
	"github.com/joseluis8906/yummies/go-code/internal/home"
	"github.com/joseluis8906/yummies/go-code/internal/menu"
)

// Module exports the module for app.
var Module = fx.Provide(
	config.New,
	log.New,
	nats.New,
	mongodb.New,
	//services
	home.New,
	menu.New,
)
