package app

import (
	"go.uber.org/fx"

	"github.com/joseluis8906/yummies/go-code/internal/home"
	"github.com/joseluis8906/yummies/go-code/internal/menu"

	"github.com/joseluis8906/yummies/go-code/pkg/config"
	"github.com/joseluis8906/yummies/go-code/pkg/log"
	"github.com/joseluis8906/yummies/go-code/pkg/mongodb"
	"github.com/joseluis8906/yummies/go-code/pkg/nats"
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
