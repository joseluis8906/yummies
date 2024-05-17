package app

import (
	"go.uber.org/fx"

	"github.com/joseluis8906/yummies/go-code/src/delivery/internal/app/bus"
	"github.com/joseluis8906/yummies/go-code/src/delivery/internal/app/config"
	"github.com/joseluis8906/yummies/go-code/src/delivery/internal/app/db"
	"github.com/joseluis8906/yummies/go-code/src/delivery/internal/app/log"
	// "github.com/joseluis8906/yummies/go-code/src/delivery/internal/store"
	// "github.com/joseluis8906/yummies/go-code/src/delivery/internal/storemanager"
	"github.com/joseluis8906/yummies/go-code/src/delivery/internal/home"
)

// Module exports the module for app.
var Module = fx.Provide(
	config.New,
	log.New,
	bus.New,
	db.New,
	//services
	home.New,
)
