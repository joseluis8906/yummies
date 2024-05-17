package app

import (
	"go.uber.org/fx"

	"github.com/joseluis8906/yummies/go-code/internal/app/bus"
	"github.com/joseluis8906/yummies/go-code/internal/app/config"
	"github.com/joseluis8906/yummies/go-code/internal/app/db"
	"github.com/joseluis8906/yummies/go-code/internal/app/log"
	// "github.com/joseluis8906/yummies/go-code/internal/store"
	// "github.com/joseluis8906/yummies/go-code/internal/storemanager"
	"github.com/joseluis8906/yummies/go-code/internal/home"
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
