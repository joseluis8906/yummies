package app

import (
	"go.uber.org/fx"

	"github.com/joseluis8906/go-code/src/delivery/internal/app/bus"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/config"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/log"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/nosql"
	"github.com/joseluis8906/go-code/src/delivery/internal/store"
	"github.com/joseluis8906/go-code/src/delivery/internal/storemanager"
)

// Module exports the module for app.
var Module = fx.Provide(
	config.New,
	log.New,
	bus.New,
	nosql.New,
	store.NewRepository,
	storemanager.New,
)
